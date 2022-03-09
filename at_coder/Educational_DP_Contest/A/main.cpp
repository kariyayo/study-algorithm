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
  int N;
  cin >> N;
  vector<int> hs(N);
  rep(i, N) cin >> hs[i];

  vector<int> dp(N, INF);
  dp[0] = 0;
  rep(i, N) {
    if (i+1 < N) {
      chmin(dp[i+1], dp[i] + abs(hs[i+1]-hs[i]));
    }
    if (i+2 < N) {
      chmin(dp[i+2], dp[i] + abs(hs[i+2]-hs[i]));
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
