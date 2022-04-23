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
  int N, K;
  cin >> N >> K;
  vector<string> ss(N);
  for (int i = 0; i < N; i++) cin >> ss[i];

  int ans = 0;
  for (int bit = 0; bit < (1 << N); bit++) {
    map<char, int> counts;
    for (int i = 0; i < N; i++) {
      if ((bit & (1 << i)) != 0) {
        string s = ss[i];
        for (int j = 0; j < (int)s.size(); j++) {
          char c = s[j];
          counts[c] = counts[c]+1;
        }
      }
    }
    int v = 0;
    for (auto iter = counts.begin(); iter != counts.end(); iter++) {
      int count = iter->second;
      if (count == K) {
        v++;
      }
    }
    chmax(ans, v);
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
