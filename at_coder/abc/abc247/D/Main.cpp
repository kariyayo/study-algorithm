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

struct Enq {
  int x, c;
};

int main() {
  int Q;
  cin >> Q;

  queue<Enq> enqs;
  for (int i = 0; i < Q; i++) {
    int query;
    cin >> query;
    if (query == 1) {
      int x, c;
      cin >> x >> c;
      Enq enq;
      enq.x = x;
      enq.c = c;
      enqs.push(enq);
    } else {
      int c;
      cin >> c;
      ll sum = 0;
      bool end = false;
      while (!end) {
        Enq & enq = enqs.front();
        if (enq.c >= c) {
          enq.c -= c;
          end = true;
          sum += 1ll * enq.x * c;
        } else {
          c -= enq.c;
          enqs.pop();
          sum += 1ll * enq.x * enq.c;
        }
      }
      cout << sum << endl;
    }
  }
}


template <typename T>
void printVector(vector<T> xs) {
  cout << '[';
  for (auto& x: xs) {
    cout << x << ' ';
  }
  cout << ']' << endl;
}
