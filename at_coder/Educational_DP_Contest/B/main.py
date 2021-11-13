INF = float('inf')

def main():
  N, K = map(int, input().split())
  xs = list(map(int, input().split()))

  dp = [INF] * (N)
  dp[0] = 0
  for i in range(0, N):
    for j in range(1, K+1):
      if i+j < N:
        dp[i+j] = min(dp[i+j], dp[i] + abs(xs[i+j]-xs[i]))
  print(dp[N-1])

##########

main()
