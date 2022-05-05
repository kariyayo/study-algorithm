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
  int N;
  cin >> N;
  vector<int> as(N);
  for (int i = 0; i < N; i++) cin >> as[i];

  int cnt = 0;
  for (int i = 0; i < N; i++) {
    int min_idx = i;
    for (int j = i; j < N; j++) {
      if (as[min_idx] > as[j]) {
        min_idx = j;
      }
    }
    if (min_idx != i) {
      int tmp = as[i];
      as[i] = as[min_idx];
      as[min_idx] = tmp;
      cnt++;
    }
  }
  printVector(as, "", "");
  cout << cnt << endl;
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
