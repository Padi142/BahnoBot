package commands

import (
	"bahno_bot/generic/record"
	"bahno_bot/generic/substance"
	"bahno_bot/generic/user"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"os/exec"
	"time"

	"github.com/bwmarrin/discordgo"
	//exposes "chart"
)

func ChartsCommand(name string, userUseCase user.UseCase, recordUseCase record.UseCase, substanceUseCase substance.UseCase) Command {
	substances, _ := substanceUseCase.GetSubstances()

	substanceChoices := make([]*discordgo.ApplicationCommandOptionChoice, len(substances))

	for i, sub := range substances {
		substanceChoices[i] = &discordgo.ApplicationCommandOptionChoice{
			Name:  sub.Label,
			Value: i,
		}
	}

	type TimePeriod struct {
		Value    string
		DayCount int
	}

	daysInWeek := 7
	daysInMonth := 30
	timePeriods := [7]TimePeriod{
		{Value: "1 tyden", DayCount: daysInWeek},
		{Value: "2 tydny", DayCount: 2 * daysInWeek},
		{Value: "3 tydny", DayCount: 3 * daysInWeek},
		{Value: "4 tydny", DayCount: 4 * daysInWeek},
		{Value: "1 mesic", DayCount: daysInMonth},
		{Value: "2 mesice", DayCount: 2 * daysInMonth},
		{Value: "3 mesice", DayCount: 3 * daysInMonth},
	}
	timePeriodsChoices := make([]*discordgo.ApplicationCommandOptionChoice, len(timePeriods))

	for i, tp := range timePeriods {
		timePeriodsChoices[i] = &discordgo.ApplicationCommandOptionChoice{
			Name:  tp.Value,
			Value: i,
		}
	}

	command := discordgo.ApplicationCommand{
		Name:        name,
		Description: "Tvuj ucet bahnaka",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "substance",
				Description: "Vyber si jakou substanci chceš grafovat",
				Choices:     substanceChoices,
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "time_period",
				Description: "Vyber si jake casove obdobi chces grafovat",
				Choices:     timePeriodsChoices,
				Required:    false,
			},
		},
	}

	handler := func(
		s *discordgo.Session,
		i *discordgo.InteractionCreate,
	) {
		//Only handle this command
		if i.ApplicationCommandData().Name != command.Name {
			return
		}
		LogCommandUse(i.Member.User.Username, command.Name)

		usr, err := userUseCase.GetProfileByDiscordID(i.Member.User.ID)
		if err != nil {
			err = SendInteractionResponse(s, i, err.Error())
			return
		}
		if err != nil {
			return
		}

		options := i.ApplicationCommandData().Options
		chosenSubstanceName := ""

		optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, opt := range options {
			optionMap[opt.Name] = opt
		}

		if opt, ok := optionMap["substance"]; ok {
			chosenSubstanceName = substances[opt.IntValue()].Value
		}

		timePeriodDays := daysInMonth

		if opt, ok := optionMap["time_period"]; ok {
			timePeriodDays = timePeriods[opt.IntValue()].DayCount
		}

		data := ChartsData{
			Day:     make([]string, timePeriodDays),
			Columns: make(map[string][]float64),
		}

		days := generateLastNDays(timePeriodDays)
		for i, day := range days {
			data.Day[i] = day.Format("2006-01-02")
		}

		if len(chosenSubstanceName) == 0 {
			for _, substance := range substances {
				data.Columns[substance.Value] = make([]float64, timePeriodDays)
			}
		} else {
			data.Columns[chosenSubstanceName] = make([]float64, timePeriodDays)
		}

		records, err := recordUseCase.GetAllRecords(usr.ID)

		for _, record := range records {

			if len(chosenSubstanceName) != 0 && record.Substance.Value != chosenSubstanceName {
				continue
			}

			dayDiff := int(math.Abs(record.CreatedAt.Sub(days[0]).Hours() / 24))

			// Out of time period
			if dayDiff >= timePeriodDays {
				continue
			}

			data.Columns[record.Substance.Value][dayDiff] += record.Amount
		}

		d, _ := json.Marshal(data)

		cmd := exec.Command("python3", "./charts.py", string(d))
		out, err := cmd.CombinedOutput()

		fmt.Println(string(out))

		// // Path to the image file
		imagePath := "output.png"

		// // Read the image file
		imageFile, err := os.Open(imagePath)
		if err != nil {
			fmt.Println("Error reading image file:", err)
			return
		}
		defer imageFile.Close()

		// Create a new file message
		file := &discordgo.File{
			Name:   "hovnohovno.png",
			Reader: imageFile,
		}

		s.ChannelMessageSendComplex(i.ChannelID, &discordgo.MessageSend{
			Files: []*discordgo.File{file},
		})

		err = SendInteractionResponse(s, i, "🤓 👉🏿 📈")

	}

	return Command{Command: command, Handler: handler}
}

func generateLastNDays(n int) []time.Time {
	days := make([]time.Time, n)

	// Generate the last 30 days
	for i := 0; i < n; i++ {
		days[i] = time.Now().Add(time.Duration(-(n - 1 - i)) * time.Hour * 24)
	}

	return days
}

type ChartsData struct {
	Day     []string
	Columns map[string][]float64
}
