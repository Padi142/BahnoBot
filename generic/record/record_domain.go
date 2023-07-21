package record

import (
	"bahno_bot/generic/models"
	"time"
)

type RecordRepository interface {
	Create(record models.Record) error
	GetAll(userId uint) (records []models.Record, err error)
	GetAllPaged(userId uint, page, pageSize int) (records []models.Record, count int64, err error)
	GetAllPagedForSubstance(userId, substanceId uint, page, pageSize int) (records []models.Record, count int64, err error)
	GetLast(userId uint) (records models.Record, err error)
	GetLastForSubstance(substanceId, userId uint) (records models.Record, err error)
	GetAllInTimePeriod(userId uint, time time.Time) (records []models.Record, err error)
	GetLeaderboardInTimePeriod(time time.Time) (leaderboard []models.LeaderboardOccurrence, err error)
	GetLeaderboardForSubstanceInTimePeriod(substanceId uint, time time.Time) (leaderboard []models.LeaderboardOccurrence, err error)
}
