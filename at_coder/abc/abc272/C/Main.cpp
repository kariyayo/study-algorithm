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
  vector<ll> odds;
  vector<ll> evens;
  for (int i = 0; i < N; i++) {
    int x;
    cin >> x;
    if (x % 2 == 0) {
      evens.push_back(x);
    } else {
      odds.push_back(x);
    }
  }
  sort(odds.begin(), odds.end(), greater<int>());
  sort(evens.begin(), evens.end(), greater<int>());
  ll ans = -1;
  if (odds.size() >= 2 && (odds[0] + odds[1]) % 2 == 0) {
    chmax(ans, odds[0] + odds[1]);
  }
  if (evens.size() >= 2 && (evens[0] + evens[1]) % 2 == 0) {
    chmax(ans, evens[0] + evens[1]);
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
