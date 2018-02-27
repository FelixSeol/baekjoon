N = int(input())

def Solve(str) :
    cnt = 0
    lstr = list(str)
    for i in lstr :
        if i == '(' :
            cnt = cnt + 1
        elif i == ')' :
            cnt = cnt - 1
        else :
            print("out of condition")
        if cnt < 0 :
            print("NO")
            return
    if cnt == 0 :
        print("YES")
    else :
        print("NO")
    
for i in range(N) :
    Solve(input())

