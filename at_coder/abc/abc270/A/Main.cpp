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
  int A, B;
  cin >> A >> B;
  int ans = 0;
  if (A == 7 || B == 7) {
    ans = 7;
  } else if (A == 6 || B == 6) {
    ans = 6;
    if (A == 1 || A == 3 || A == 5) {
      ans += 1;
    } else if (B == 1 || B == 3 || B == 5) {
      ans += 1;
    }
  } else if (A == 5 || B == 5) {
    ans = 5;
    if (A == 2 || A == 3) {
      ans += 2;
    } else if (B == 2 || B == 3) {
      ans += 2;
    }
  } else if (A == 3 || B == 3) {
    ans = 3;
    if (A == 4 || B == 4) {
      ans += 4;
    }
  } else {
    if (A == B) {
      ans += A;
    } else {
      ans += A;
      ans += B;
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
