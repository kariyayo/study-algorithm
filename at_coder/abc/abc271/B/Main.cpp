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
  int N, Q;
  cin >> N >> Q;

  vector<int> L(N);
  vector<vector<int>> As(N);
  for (int i = 0; i < N; i++) {
    int l;
    cin >> l;
    L[i] = l;
    vector<int> A(l);
    for (int j = 0; j < l; j++) {
      cin >> A[j];
    }
    As[i] = A;
  }
  for (int k = 0; k < Q; k++) {
    int s, t;
    cin >> s >> t;
    cout << As[s-1][t-1] << endl;
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
