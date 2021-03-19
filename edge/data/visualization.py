import seaborn
import matplotlib.pyplot as plt
import pandas as pd
import os
import numpy as np
converter = {'ball': 0, 'inner': 0.01,
             'comb': 0.02, 'outer': 0.03, 'health': 0.04}


def converte(type):
    return converter[type]


names = ['Channel1', 'Channel2', 'Channel3', 'Channel4',
         'Channel5', 'Channel6', 'Channel7', 'Channel8', 'Classification', ]
path = os.getcwd()+"/edge/data/dataset"
df = pd.read_csv(path+"/data.csv", header=None, engine='c',
                 na_filter=False, names=names, converters={'Classification': converte})
# df = pd.read_csv(path+"/data.csv", header=None, engine='c',
#                  na_filter=False, names=names)
print(path, "read done")
print(df)
df_corr = df.corr()
seaborn.heatmap(df_corr, center=0, annot=True, cmap='YlGnBu')
plt.show()

df_cov = df.cov()
seaborn.heatmap(df_cov, center=0, annot=True,)
plt.show()

# fig, ax = plt.subplots(3, 3)
# df = df.sample(n=20000, axis=0)
# cmap = plt.cm.get_cmap("hsv", 9)
# for i in range(3):
#     for j in range(3):
#         name = names[i*3+j]
#         if name != 'Classification':
#             ax[i][j].scatter(df['Classification'], df[name],  c=cmap(i*3+j),
#                              label=name, alpha=0.3, edgecolors='none')
#             ax[i][j].legend()
#             ax[i][j].grid(True)
#             ax[i][j].set_xlabel('Classification')
#             ax[i][j].set_ylabel(name)

# plt.tight_layout()
# plt.show()
