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
    plt.tick_params(axis='x', colors='grey')
    plt.tick_params(axis='y', colors='grey')

    plt.xlabel('X-axis Label', color='grey')
    plt.ylabel('amount [g]', color='grey')
    plt.legend(loc='upper center', bbox_to_anchor=(0.5, 1.15), ncol=2)

    plt.savefig("output.png", transparent=True)

if __name__ == "__main__":
    if len(sys.argv) != 2:
        exit(1)

    data = sys.argv[1]

    generate_chart(json.loads(data))