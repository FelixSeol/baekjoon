num = input()
mList = list()

for i in range(int(num)) :
    mList.append(input())
result = mList[0]

length = len(mList[0])
for j in range(length) :
    for i in range(int(num)-1) :
        if mList[i][j] != mList[i+1][j] :
            result = result[:j] + '?' + result[j+1:]
            break

print(result)



