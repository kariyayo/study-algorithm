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
  int X, Y, Z;
  cin >> X >> Y >> Z;

  if ((X > 0 && Y < 0) || (X > 0 && Y > X)) {
    cout << X << endl;
  } else if ((X < 0 && Y > 0) || (X < 0 && Y < X)) {
    cout << -X << endl;
  } else if (X > 0 && Y > 0) {
    if (Z < 0) {
      cout << (-Z) + (-Z) + X << endl;
    } else if (Z > Y) {
      cout << -1 << endl;
    } else {
      cout << X << endl;
    }
  } else if (X < 0 && Y < 0) {
    if (Z > 0) {
      cout << (Z) + (Z) + (-X) << endl;
    } else if (Z < Y) {
      cout << -1 << endl;
    } else {
      cout << -X << endl;
    }
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
