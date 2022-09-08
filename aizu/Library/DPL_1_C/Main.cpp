#include <bits/stdc++.h>

using namespace std;
using uint = unsigned int;
using ll = long long;
using ull = unsigned long long;
template <typename T> inline bool chmin(T& a, T b) { if (a > b) { a = b; return true; } return false; }
template <typename T> inline bool chmax(T& a, T b) { if (a < b) { a = b; return true; } return false; }
template <typename T> void printVector(vector<T> xs, string prefix = "[ ", string suffix = " ]");

const long long INF_LL = 1LL<<60;
const int INF = 1000010000;


int main() {
  int N, W;
  cin >> N >> W;
  vector<int> vs(N+1);
  vector<int> ws(N+1);
  for (int i = 1; i <= N; i++) {
    cin >> vs[i] >> ws[i];
  }
  vector<vector<int>> dp(N+1);
  dp[0] = vector<int>(W+1);
  for (int i = 1; i <= N; i++) {
    dp[i] = vector<int>(W+1);
    for (int w = 1; w <= W; w++) {
      if (w - ws[i] >= 0) {
        // 1つ追加
        chmax(dp[i][w], dp[i][w-ws[i]] + vs[i]);
        // 追加しない
        chmax(dp[i][w], dp[i-1][w]);
      } else {
        // 追加する余地がない
        chmax(dp[i][w], dp[i-1][w]);
      }
    }
  }
  int ans = 0;
  for (int w = 1; w <= W; w++) {
    chmax(ans, dp[N][w]);
  }
  cout << ans << endl;
}

template <typename T>
void printVector(vector<T> xs, string prefix, string suffix) {
  cout << prefix;
  for (int i = 0; i < (ll)xs.size(); i++) {
    cout << xs[i];
    if (i+1 != (ll)xs.size()) {
      cout << " ";
    }
  }
  cout << suffix << endl;
}
