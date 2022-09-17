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
  vector<string> ss(10);
  int a = 0;
  int b = 0;
  int c = 0;
  int d = 0;
  for (int i = 0; i < 10; i++) {
    string s;
    cin >> s;
    for (int j = 0; j < 10; j++) {
      if (s[j] == '#') {
        if (a == 0) {
          a = i+1;
        }
        if (c == 0) {
          c = j+1;
        }
      }
      if (s[j] == '.') {
        if (c != 0 && d == 0) {
          d = j;
        }
        if (d != 0 && c == j+1 && b == 0) {
          b = i;
        }
      }
    }
    if (c != 0 && d == 0) d = 10;
  }
  if (a != 0 && b == 0) b = 10;
  cout << a << " " << b << endl << c << " " << d << endl;
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
