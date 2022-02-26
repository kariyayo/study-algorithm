#include <bits/stdc++.h>
#include <atcoder/all>

using namespace std;

using uint = unsigned int;
using ll = long long;
using ull = unsigned long long;
const ll MOD = 1000000007LL;
const int INF = 20002000;

#define rep(i, n) for (int i = 0; i < (int)(n); ++i)
#define rep2(i, start, n) for (int i = (start); i < (int)(n); ++i)

int main() {
  int K, S;
  cin >> K >> S;
  int ans = 0;
  rep(x, K+1) {
    rep(y, K+1) {
      int z = S - (x + y);
      if (z >= 0 && z <= K) {
        ans++;
      }
    }
  }
  cout << ans << endl;
}
