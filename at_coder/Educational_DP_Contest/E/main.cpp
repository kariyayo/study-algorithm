#include <bits/stdc++.h>
#include <atcoder/all>

using namespace std;

using uint = unsigned int;
using ll = long long;
using ull = unsigned long long;

#define rep(i, n) for (int i = 0; i < (int)(n); ++i)
#define rep2(i, start, n) for (int i = (start); i < (int)(n); ++i)

template <typename T> inline bool chmin(T& a, T b) { if (a > b) { a = b; return true; } return false; }
template <typename T> inline bool chmax(T& a, T b) { if (a < b) { a = b; return true; } return false; }
template <typename T> void printVector(vector<T> xs);

const long long INF = 1LL<<60;
// const int INF = 1000010000;

int main() {
  int N, W;
  cin >> N >> W;
  vector<int> ws(N+1), vs(N+1);
  rep(i, N) cin >> ws[i+1] >> vs[i+1];

  int v_sum_max = 1000 * 100;

  vector<vector<ll>> dp(N+1, vector<ll>(v_sum_max+1, INF));
  dp[0][0] = 0;
  rep2(i, 1, N+1) {
    rep2(vi, 0, v_sum_max+1) {
      // i番目を選ばない
      chmin(dp[i][vi], dp[i-1][vi]);

      // i番目を選ぶ
      if (vi-vs[i] >= 0) {
        chmin(dp[i][vi], ws[i] + dp[i-1][vi-vs[i]]);
      }
    }
  }
  int ans = 0;
  rep(v, v_sum_max+1) {
    ll w_sum = dp[N][v];
    if (w_sum <= W) {
      chmax(ans, v);
    }
  }
  cout << ans << endl;
}


template <typename T>
void printVector(vector<T> xs) {
  cout << '[';
  for (auto& x: xs) {
    cout << x << ' ';
  }
  cout << ']' << endl;
}
