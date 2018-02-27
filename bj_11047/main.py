N, K = input().split()
N = int(N)
K = int(K)
cnt = 0

coins = list(range(N))
for i in range(N) :
    coins[i] = int(input())

coins.reverse()


for i in coins :
    if K >= i :
        q = int(K / i)
        K = K - q*i
        cnt = cnt + q
    if K == 0 :
    	break

print(cnt)
