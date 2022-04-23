#include <bits/stdc++.h>
#include <atcoder/all>
using namespace std;
using uint = unsigned int;
using ll = long long;
using ull = unsigned long long;
template <typename T> inline bool chmin(T& a, T b) { if (a > b) { a = b; return true; } return false; }
template <typename T> inline bool chmax(T& a, T b) { if (a < b) { a = b; return true; } return false; }
template <typename T> void printVector(vector<T> xs, string prefix = "[ ", string suffix = " ]");

// const long long INF = 1LL<<60;
const int INF = 1000010000;

int main() {
  int A, B, C, D, E, F, X;
  cin >> A >> B >> C >> D >> E >> F >> X;
  int v1 = 0;
  int v2 = 0;
  int a = A;
  for (int t = 0; t < X;) {
    if (a > 0) {
      v1 += B;
      a--;
      t++;
    } else {
      a = A;
      t = t + C;
    }
  }
  int d = D;
  for (int t = 0; t < X;) {
    if (d > 0) {
      v2 += E;
      d--;
      t++;
    } else {
      d = D;
      t = t + F;
    }
  }
  if (v1 > v2) {
    cout << "Takahashi" << endl;
  } else if (v1 < v2) {
    cout << "Aoki" << endl;
  } else {
    cout << "Draw" << endl;
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
