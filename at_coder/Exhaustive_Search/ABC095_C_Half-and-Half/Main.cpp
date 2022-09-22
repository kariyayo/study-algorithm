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
  int A, B, C, X, Y;
  cin >> A >> B >> C >> X >> Y;

  ll ans = 0;
  int k = min(X, Y);
  if (A + B > 2*C) {
    ans += 2 * C * k;
  } else {
    ans += (A + B) * k;
  }
  int d = abs(X - Y);
  if (X > Y) {
    if (A < 2*C) {
      ans += A * d;
    } else {
      ans += 2 * C * d;
    }
  } else {
    if (B < 2*C) {
      ans += B * d;
    } else {
      ans += 2 * C * d;
    }
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
