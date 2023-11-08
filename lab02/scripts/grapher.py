import sys
import seaborn
import pandas as pd
from matplotlib import pyplot as plt

class GraphCreator:
    def __init__(self, *args, dst_fname='graph.png'):
        self.data_files = [arg for arg in args]
        self.dst_fname = dst_fname
        
    def read_data(self, fname):
        data = pd.read_csv(fname)
        return data
    
    def plot(self):
        for plot_fname in self.data_files:
            data = data=self.read_data(plot_fname)
            seaborn.lineplot(data=data, x='len', y=data.axes[1][1], label=data.axes[1][1], marker='o')
        plt.yscale('log')
        plt.xlabel('Количество строк квадратных матриц, шт')
        plt.ylabel('Память, байт')
        plt.grid()
        plt.savefig(self.dst_fname)
        

if __name__ == '__main__':
    if len(sys.argv) < 2:
        raise Exception("Ошибка с аргументами командной строки!")
    
    files = []
    for arg in sys.argv[1:]:
        files.append(arg) 
   
    g_creator = GraphCreator(*files)     
    g_creator.plot() 
