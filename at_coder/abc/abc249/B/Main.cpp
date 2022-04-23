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
  string S;
  cin >> S;

  set<int> ss;
  bool hasUpper = false;
  bool hasLower = false;
  for (int i = 0; i < (int)S.size(); i++) {
    char c = S[i];
    ss.insert(c);
    if (islower(c)) {
      hasLower = true;
    }
    if (isupper(c)) {
      hasUpper = true;
    }
  }
  if (S.size() == ss.size() && hasUpper && hasLower) {
    cout << "Yes" << endl;
  } else {
    cout << "No" << endl;
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
