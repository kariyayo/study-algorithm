#include <bits/stdc++.h>
#include <atcoder/all>
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
  int N, S;
  cin >> N >> S;
  vector<int> as(N);
  vector<int> bs(N);
  for (int i = 0; i < N; i++) {
    cin >> as[i] >> bs[i];
  }
  vector<vector<bool>> dp(N+1, vector<bool>(S+1));
  dp[0][0] = true;
  for (int i = 1; i <= N; i++) {
    int a = as[i - 1];
    int b = bs[i - 1];
    for (int s = 1; s <= S; s++) {
      if (s-a >= 0 && dp[i-1][s-a]) {
        dp[i][s] = true;
      }
      if (s-b >= 0 && dp[i-1][s-b]) {
        dp[i][s] = true;
      }
    }
  }
  if (dp[N][S]) {
    cout << "Yes" << '\n';
    string ans(N, 'H'); // 初期値は全て表
    int s = S;
    for (int i = N - 1; i >= 0; i--) {
      int a = as[i];
      int b = bs[i];
      if (s >= a && dp[i][s-a]) {
        ans[i] = 'H'; // 表
        s -= a;
      } else {
        ans[i] = 'T'; // 裏
        s -= b;
      }
    }
    cout << ans << '\n';
    return 0;
  }
  cout << "No" << '\n';
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
