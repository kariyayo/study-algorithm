def main():
  N = int(input())
  xs = list(map(int, input().split()))

  dp = [101010100] * N
  dp[0] = 0
  for i in range(1, N):
    dp[i] = min(dp[i], dp[i-1] + abs(xs[i]-xs[i-1]))
    if (i > 1):
      dp[i] = min(dp[i], dp[i-2] + abs(xs[i]-xs[i-2]))
  print(dp[N-1])

##########

main()
