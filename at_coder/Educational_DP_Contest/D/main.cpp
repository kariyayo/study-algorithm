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
  int N, W;
  cin >> N >> W;
  vector<int> ws(N+1), vs(N+1);
  rep(i, N) cin >> ws[i+1] >> vs[i+1];

  vector<vector<ll>> dp(N+1, vector<ll>(W+1));
  dp[0][0] = 0;

  rep2(i, 1, N+1) {
    rep2(w, 1, W+1) {
      // i番目を選ばない
      chmax(dp[i][w], dp[i-1][w]);
      // i番目を選ぶ
      if (w-ws[i] >= 0) {
        chmax(dp[i][w], vs[i] + dp[i-1][w-ws[i]]);
      }
    }
  }
  cout << dp[N][W] << endl;
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
