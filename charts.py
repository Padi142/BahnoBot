import seaborn as sns
import matplotlib.pyplot as plt
import matplotlib.ticker as ticker
import matplotlib.dates as mdates
import pandas as pd
import sys
from typing import Dict, List
import json
import datetime
  
def generate_chart(data: Dict):
    d = data["Columns"]
    d["day"] = data["Day"]

    df = pd.DataFrame(d) 

    plt.figure(figsize=(7, 6))

    sns.lineplot(x='day', y='value', hue='variable', 
                data=pd.melt(df, ['day']),
                 linewidth=2.3)

    ticks = plt.xticks()[0]
    labels = plt.xticks()[1]

    # Keep only every 4th tick
    new_ticks = ticks[::4]
    new_labels = labels[::4]
    plt.xticks(new_ticks, new_labels)

    # Rotate the date tick labels for better visibility (optional)
    plt.xticks(rotation=45)
    plt.tick_params(axis='x', colors='white')
    plt.tick_params(axis='y', colors='white')

    plt.xlabel('', color='white', labelpad=10)
    plt.ylabel('amount [g]', color='white', labelpad=10)
    plt.legend(loc='upper center', bbox_to_anchor=(0.5, 1.15), ncol=2)

    plt.subplots_adjust(bottom=0.2)

    plt.savefig("output.png", transparent=True)

if __name__ == "__main__":
    if len(sys.argv) != 2:
        exit(1)

    data = sys.argv[1]

    generate_chart(json.loads(data))