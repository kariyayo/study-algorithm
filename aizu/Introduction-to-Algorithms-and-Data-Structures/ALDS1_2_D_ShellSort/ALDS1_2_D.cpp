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

int cnt = 0;

void insertionSort(vector<int> &as, int n, int g) {
  for (int i = g; i < n; i++) {
    int v = as[i];
    int j = i - g; // j番目までソート済みにする
    while (j >= 0 && as[j] > v) {
      as[j+g] = as[j];
      j = j - g;
      cnt++;
    }
    as[j+g] = v;
  }
}

void shellSort(vector<int> as, int n) {
  vector<int> G;
  for (int h = 1; ; ) {
    if (h > n) break;
    G.push_back(h);
    h = 3*h + 1;
  }
  int m = G.size();
  cout << m << endl;
  for (int i = m-1; i >= 0; i--) {
    insertionSort(as, n, G[i]);
    cout << G[i];
    if (i != 0) cout << " ";
  }
  cout << endl;
  cout << cnt << endl;
  for (int i = 0; i < n; i++) {
    cout << as[i] << endl;
  }
}

int main() {
  int n;
  cin >> n;
  vector<int> as(n);
  for (int i = 0; i < n; i++) cin >> as[i];

  shellSort(as, n);
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
