#include <bits/stdc++.h>
#include <atcoder/all>

using namespace std;

using uint = unsigned int;
using ll = long long;
using ull = unsigned long long;
const int INF = 1000010000;

#define rep(i, n) for (int i = 0; i < (int)(n); ++i)
#define rep2(i, start, n) for (int i = (start); i < (int)(n); ++i)

template <typename T> void chmin(T& a, T b);
template <typename T> void chmax(T& a, T b);
template <typename T> void printVector(vector<T> xs);

int main() {
  int N;
  cin >> N;
  vector<int> as(N+1);
  vector<int> bs(N+1);
  vector<int> cs(N+1);
  rep2(i, 1, N+1) {
    cin >> as[i] >> bs[i] >> cs[i];
  }

  vector<int> dp_a(N+1);
  vector<int> dp_b(N+1);
  vector<int> dp_c(N+1);

  rep2(i, 1, N+1) {
    chmax(dp_a[i], as[i] + dp_b[i-1]);
    chmax(dp_a[i], as[i] + dp_c[i-1]);

    chmax(dp_b[i], bs[i] + dp_a[i-1]);
    chmax(dp_b[i], bs[i] + dp_c[i-1]);

    chmax(dp_c[i], cs[i] + dp_a[i-1]);
    chmax(dp_c[i], cs[i] + dp_b[i-1]);
  }

  cout << max(max(dp_a[N], dp_b[N]), dp_c[N]) << endl;
}


template <typename T>
void chmin(T& a, T b) {
  if (a > b) {
    a = b;
  }
}

template <typename T>
void chmax(T& a, T b) {
  if (a < b) {
    a = b;
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
