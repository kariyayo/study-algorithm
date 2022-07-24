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
  int K;
  cin >> K;

  int h_add, m_add;
  if (K >= 60) {
    h_add = 1;
    m_add = K - 60;
  } else {
    h_add = 0;
    m_add = K;
  }

  int hh = 21 + h_add;
  if (hh >= 24) {
    hh -= 24;
  }
  int mm = m_add;
  cout << setw(2) << setfill('0') << hh << ":" << setw(2) << setfill('0') << mm << endl;
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
