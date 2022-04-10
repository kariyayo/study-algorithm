#include <bits/stdc++.h>
#include <atcoder/all>
using namespace std;
using uint = unsigned int;
using ll = long long;
using ull = unsigned long long;
template <typename T> inline bool chmin(T& a, T b) { if (a > b) { a = b; return true; } return false; }
template <typename T> inline bool chmax(T& a, T b) { if (a < b) { a = b; return true; } return false; }
template <typename T> void printVector(vector<T> xs);

// const long long INF = 1LL<<60;
const int INF = 1000010000;

int main() {
  int N;
  cin >> N;
  vector<string> ss(N), ts(N);
  for (int i = 0; i < N; i++) cin >> ss[i] >> ts[i];

  for (int i = 0; i < N; i++) {
    // i番目の人のあだ名を考える
    string a1 = ss[i];
    string a2 = ts[i];
    bool a1OK = true, a2OK = true;
    for (int j = 0; j < N; j++) {
      if (i != j) {
        if (a1 == ss[j] || a1 == ts[j]) {
          // a1は使えない
          a1OK = false;
        }
        if (a2 == ss[j] || a2 == ts[j]) {
          // a1は使えない
          a2OK = false;
        }
      }
    }
    if (!a1OK && !a2OK) {
      // 性名どちらのあだ名もダメ
      cout << "No" << endl;
      return 0;
    }
  }
  cout << "Yes" << endl;
}


template <typename T>
void printVector(vector<T> xs) {
  cout << '[';
  for (auto& x: xs) {
    cout << x << ' ';
  }
  cout << ']' << endl;
}
