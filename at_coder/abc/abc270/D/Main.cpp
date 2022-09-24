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


vector<int> A;
int K;
vector<vector<int>> memo;

/**
 * @param player 0: 高橋くん 1: 青木くん
 * @param n 残りの石の個数
 * @return int 石が残りn個でplayerの番のときの、高橋くんの得点
 */
int solve(int player, int n) {
  if (memo[player][n] != -1) {
    return memo[player][n];
  }
  if (player == 0) {
    // 高橋くん
    int ans = 0;
    for (int i = 0; i < K; i++) {
      if (n - A[i] >= 0) {
        int AOKI_Kun = 1;
        chmax(ans, A[i] + solve(AOKI_Kun, n - A[i]));
      }
    }
    memo[player][n] = ans;
    return ans;
  } else {
    // 青木くん
    // 自分が最大値を取る→高橋くんが最小となるような手を選ぶ
    int ans = INF;
    for (int i = 0; i < K; i++) {
      if (n - A[i] >= 0) {
        int TAKAHASHI_Kun = 0;
        chmin(ans, solve(TAKAHASHI_Kun, n - A[i]));
      }
    }
    if (ans == INF) {
      ans = 0;
    }
    memo[player][n] = ans;
    return ans;
  }
}

int main() {
  int N;
  cin >> N >> K;
  A = vector<int>(K);
  for (int i = 0; i < K; i++) {
    cin >> A[i];
  }
  memo = vector<vector<int>>(2, vector<int>(N+1, -1));
  solve(0, N);
  cout << memo[0][N] << endl;
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
