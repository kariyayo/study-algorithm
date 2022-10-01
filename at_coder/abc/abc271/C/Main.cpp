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
  int N;
  cin >> N;
  vector<int> as(N);
  for (int i = 0; i < N; i++) {
    cin >> as[i];
  }
  set<int> s(as.begin(), as.end());
  vector<int> ss(s.begin(), s.end());
  sort(ss.begin(), ss.end());
  int ans = 0;
  int count = N;
  int i = 0;
  while (i < count) {
    if (ans+1 == ss[i]) {
      // 持ってた
      ans++;
      i++;
    } else {
      // 持ってなかった
      count -= 2;
      if (i > count) {
        break;
      }
      ans++;
    }
  }
  cout << ans << endl;
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
