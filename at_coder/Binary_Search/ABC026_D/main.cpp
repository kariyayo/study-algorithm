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

int A, B, C;

double f(double t) {
  return A*t + B*sin(M_PI * C * t);
}

int main() {
  cin >> A >> B >> C;
  double left = 0;
  double right = 200;
  double ans = 0.0;
  while (right - left > 0) {
    double mid = (left + right) / 2.0;
    double d = 100.0 - f(mid);
    if (abs(d) <= 0.000001) {
      ans = mid;
      break;
    } else if (d > 0) {
      left = mid;
    } else {
      right = mid;
    }
  }
  cout << fixed << setprecision(20) << ans << endl;
}


template <typename T>
void printVector(vector<T> xs) {
  cout << '[';
  for (auto& x: xs) {
    cout << x << ' ';
  }
  cout << ']' << endl;
}
