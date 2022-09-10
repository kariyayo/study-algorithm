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
  int n, m;
  cin >> n >> m;
  vector<int> cs(m);
  for (int i = 0; i < m; i++) {
    cin >> cs[i];
  }

  vector<vector<int>> dp(m, vector<int>(n+1, INF));
  for (int i = 0; i <= n; i++) {
    dp[0][i] = i;
  }

  for (int i = 1; i < m; i++) {
    for (int j = 0; j <= n; j++) {
      // 選ばない
      chmin(dp[i][j], dp[i-1][j]);

      // 選ぶ
      if (j-cs[i] >= 0) {
        chmin(dp[i][j], dp[i][j-cs[i]] + 1);
      }
    }
  }
  cout << dp[m-1][n] << endl;
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
