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

// const long long INF = 1LL<<60;
const int INF = 1000010000;

int main() {
  int N;
  cin >> N;
  vector<int> as(N), bs(N), cs(N);
  rep(i, N) cin >> as[i];
  rep(i, N) cin >> bs[i];
  rep(i, N) cin >> cs[i];

  sort(as.begin(), as.end());
  sort(cs.begin(), cs.end());
  ll ans = 0;
  rep(i, N) {
    int b = bs[i];
    auto ai_iter = lower_bound(as.begin(), as.end(), b);
    size_t ai = distance(as.begin(), ai_iter);
    auto ci_iter = upper_bound(cs.begin(), cs.end(), b);
    size_t ci = distance(cs.begin(), ci_iter);
    ll cnt = ai * (N-ci);
    ans += cnt;
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
