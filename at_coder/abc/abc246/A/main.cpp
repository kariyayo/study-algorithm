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
  int x1, y1, x2, y2, x3, y3;
  cin >> x1 >> y1;
  cin >> x2 >> y2;
  cin >> x3 >> y3;

  int x = 0;
  if (x1 == x2) {
    x = x3;
  } else if (x2 == x3) {
    x = x1;
  } else {
    x = x2;
  }
  int y = 0;
  if (y1 == y2) {
    y = y3;
  } else if (y2 == y3) {
    y = y1;
  } else {
    y = y2;
  }
  cout << x << " " << y << endl;
}


template <typename T>
void printVector(vector<T> xs) {
  cout << '[';
  for (auto& x: xs) {
    cout << x << ' ';
  }
  cout << ']' << endl;
}
