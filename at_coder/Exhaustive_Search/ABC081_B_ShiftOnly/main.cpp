#include <bits/stdc++.h>
#include <atcoder/all>

using namespace std;

using uint = unsigned int;
using ll = long long;
using ull = unsigned long long;
const ll MOD = 1000000007LL;
const int INF = 20002000;

#define rep(i, n) for (int i = 0; i < (int)(n); ++i)
#define rep2(i, s, n) for (int i = (s); i < (int)(n); ++i)

int main() {
  int N;
  cin >> N;
  vector<int> a(N);
  rep(i, N) cin >> a[i];

  int ans = 0;
  while (true) {
    rep(i, N) {
      if (a[i] < 2 || a[i] % 2 != 0) {
        cout << ans << endl;
        return 0;
      }
      a[i] = a[i] / 2;
    }
    ans++;
  }
}
