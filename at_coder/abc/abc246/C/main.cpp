#include <bits/stdc++.h>
#include <atcoder/all>
using namespace std;
using uint = unsigned int;
using ll = long long;
using ull = unsigned long long;
template <typename T> inline bool chmin(T& a, T b) { if (a > b) { a = b; return true; } return false; }
template <typename T> inline bool chmax(T& a, T b) { if (a < b) { a = b; return true; } return false; }
template <typename T> void printVector(vector<T> xs);

// const long long INF = 1LL<<60;
const int INF = 1000010000;

int main() {
  int N, K, X;
  cin >> N >> K >> X;
  vector<int> as(N);
  for (int i = 0; i < N; ++i) cin >> as[i];

  // クーポン金額より高い分を引く
  sort(as.begin(), as.end(), greater<int>());
  for (int i = 0; i < N; ++i) {
    int m = as[i] / X; // m枚までならマイナスにならない
    if (m > K) {
      m = K;
    }
    K = K - m;
    as[i] = as[i] - (X * m);
    if (K == 0) break;
  }

  if (K > 0) {
    // まだクーポンが残ってる場合、できるだけ無駄にしないように残額が高い方から引く
    sort(as.begin(), as.end(), greater<int>());
    int i = 0;
    while (K > 0 && i < N) {
      if (as[i] > 0) {
        --K;
        // 金額引く
        as[i] = max(as[i] - X, 0);
      }
      ++i;
    }
  }

  ll ans = 0;
  for (int i = 0; i < N; ++i) {
    ans += as[i];
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
