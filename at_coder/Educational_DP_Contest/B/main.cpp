#include <bits/stdc++.h>
#include <atcoder/all>

using namespace std;

using uint = unsigned int;
using ll = long long;
using ull = unsigned long long;
const int INF = 1000010000;

#define rep(i, n) for (int i = 0; i < (int)(n); ++i)
#define rep2(i, start, n) for (int i = (start); i < (int)(n); ++i)

template <typename T> void chmin(T& a, T b);
template <typename T> void chmax(T& a, T b);
template <typename T> void printVector(vector<T> xs);

int main() {
  int N, K;
  cin >> N >> K;
  vector<int> hs(N);
  rep(i, N) cin >> hs[i];

  vector<int> dp(N, INF);
  dp[0] = 0;

  rep(i, N) {
    rep2(j, 1, K+1) {
      if (i+j < N) {
        chmin(dp[i+j], dp[i] + abs(hs[i+j] - hs[i]));
      }
    }
  }
  cout << dp[N-1] << endl;
}


template <typename T>
void chmin(T& a, T b) {
  if (a > b) {
    a = b;
  }
}

template <typename T>
void chmax(T& a, T b) {
  if (a < b) {
    a = b;
  }
}

template <typename T>
void printVector(vector<T> xs) {
  cout << '[';
  for (auto& x: xs) {
    cout << x << ' ';
  }
  cout << ']' << endl;
}
