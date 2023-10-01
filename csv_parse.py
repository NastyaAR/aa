N = 7

files = [open(f"/home/nastya/Рабочий стол/aa/lab01/report/inc/csv/{i}_300.csv", "w") for i in range(4)]

for i in range(len(files)):
    files[i].write(f'len,{i+1}' + "\n")

with open("/home/nastya/Рабочий стол/aa/lab01/report/inc/csv/time.csv", "r") as src:
    title = src.readline()
    lines = []
    for i in range(N):
        line = src.readline().rstrip('\n').split(',')
        for i in range(len(files)):
            files[i].write(line[0] + "," +  line[i+1] + '\n')
        
for f in files:
    f.close()
