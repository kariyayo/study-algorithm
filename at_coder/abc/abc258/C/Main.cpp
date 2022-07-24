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


// ABC 258

int main() {
  int N, Q;
  cin >> N >> Q;
  string S;
  cin >> S;
  int start = 0;
  for (int i = 0; i < Q; i++) {
    int q, x;
    cin >> q >> x;
    if (q == 1) {
      start = (N + start - x) % N;
    } else {
      x = x - 1;
      cout << S[(start + x) % N] << '\n';
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
