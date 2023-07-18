import os

datas= set()
with open('proxy/2.html', 'r') as f:
    for line in f:
        print(line.strip())
        datas.add(line)

writeFile = 'proxy/3.txt'
if not os.path.exists(writeFile):
    os.system(r'touch {}'.format(writeFile))

with open(writeFile, 'a', encoding="utf-8") as f:
    f.writelines(datas)
print(datas)
