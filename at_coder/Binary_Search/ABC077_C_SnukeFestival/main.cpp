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
  int N;
  cin >> N;
  vector<int> as(N);
  vector<int> bs(N);
  vector<int> cs(N);
  for (int i = 0; i < N; i++) cin >> as[i];
  for (int i = 0; i < N; i++) cin >> bs[i];
  for (int i = 0; i < N; i++) cin >> cs[i];
  sort(as.begin(), as.end());
  sort(cs.begin(), cs.end());

  ll ans = 0;
  for (int i = 0; i < N; i++) {
    int b = bs[i];
    ll aCnt = lower_bound(as.begin(), as.end(), b) - as.begin();
    ll cCnt = cs.end() - upper_bound(cs.begin(), cs.end(), b);
    ans += aCnt * cCnt;
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
