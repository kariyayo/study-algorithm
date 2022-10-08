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
  int N, M;
  cin >> N >> M;
  vector<vector<int>> xss(M);
  for (int i = 0; i < M; i++) {
    int k;
    cin >> k;
    vector<int> xs(k);
    for (int j = 0; j < k; j++) {
      cin >> xs[j];
    }
    xss[i] = xs;
  }
  for (int i = 1; i < N; i++) {
    for (int j = i+1; j <= N; j++) {
      bool ok = false;
      for (int k = 0; k < M; k++) {
        vector<int> xs = xss[k];
        if (count(xs.begin(), xs.end(), i) > 0 && count(xs.begin(), xs.end(), j) > 0) {
          ok = true;
          break;
        }
      }
      if (!ok) {
        cout << "No" << endl;
        return 0;
      }
    }
  }
  cout << "Yes" << endl;
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
