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
  ll N;
  cin >> N;
  vector<ll> res = { 0 };
  for (int i = 0; i < 60; i++) {
    if ((N & (1ll << i)) != 0) {
      int sz = res.size();
      for (int j = 0; j < sz; j++) {
        res.push_back(res[j] | (1ll << i));
      }
    }
  }
  for (int i = 0; i < (int)res.size(); i++) {
    cout << res[i] << endl;
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
