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
  int q;
  cin >> q;
  for (int i = 0; i < q; i++) {
    string X, Y;
    cin >> X >> Y;
    vector<vector<int>> dp(X.length()+1, vector<int>(Y.length()+1));
    for (int i = 1; i <= (int)X.length(); i++) {
      for (int j = 1; j <= (int)Y.length(); j++) {
        if (X[i-1] == Y[j-1]) {
          chmax(dp[i][j], dp[i-1][j-1] + 1); // マッチして+1
        }
        chmax(dp[i][j], dp[i-1][j]); // マッチしたけど+0 or マッチせず -> X[i]を捨ててY[j]を生かす
        chmax(dp[i][j], dp[i][j-1]); // マッチせず -> X[i]を生かしてY[j]を捨てる
      }
    }
    cout << dp[X.length()][Y.length()] << endl;
  }
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
