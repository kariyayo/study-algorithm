#include <bits/stdc++.h>
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
    for (int j = N - 1; j > i; j--) {
      if (as[j-1] > as[j]) {
        int tmp = as[j];
        as[j] = as[j-1];
        as[j-1] = tmp;
        cnt++;
      }
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
