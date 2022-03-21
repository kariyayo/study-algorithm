#include <bits/stdc++.h>
#include <atcoder/all>

using namespace std;

using uint = unsigned int;
using ll = long long;
using ull = unsigned long long;

#define rep(i, n) for (int i = 0; i < (int)(n); ++i)
#define rep2(i, start, n) for (int i = (start); i < (int)(n); ++i)

template <typename T> inline bool chmin(T& a, T b) { if (a > b) { a = b; return true; } return false; }
template <typename T> inline bool chmax(T& a, T b) { if (a < b) { a = b; return true; } return false; }
template <typename T> void printVector(vector<T> xs);

// const long long INF = 1LL<<60;
const int INF = 1000010000;

int main() {
  int N;
  cin >> N;

  vector<int> hs(N), ss(N);
  rep(i, N) {
    cin >> hs[i] >> ss[i];
  }

  ll left = 0L;
  ll right = 1000000000L + 1000000000L * 100000L; // H_MAX + S_MAX * N_MAX
  while (right - left > 1) {
    ll mid = (right + left) / 2;

    // midが条件を満たしているか？ => 最大mid点で風船を割れるか？
    bool ok = true;

    // 高度midに達するまでの時間
    vector<ll> ts(N);
    rep(i, N) {
      if (mid - hs[i] < 0) {
        ok = false;
      } else {
        ts[i] = (mid - hs[i]) / ss[i];
      }
    }

    // 早く割らないといけない順に並べ替え
    sort(ts.begin(), ts.end());

    // 時間を1秒ずつ進める
    rep(i, N) {
      if (ts[i] < i) {
        ok = false;
      }
    }

    if (ok) {
      // もっと小さくできる
      right = mid;
    } else {
      // もっと大きくしないとダメ
      left = mid;
    }
  }

  cout << right << endl;
}


template <typename T>
void printVector(vector<T> xs) {
  cout << '[';
  for (auto& x: xs) {
    cout << x << ' ';
  }
  cout << ']' << endl;
}
