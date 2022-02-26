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
  string S;
  cin >> S;
  int N = S.size();
  ll ans = 0;
  for (int bit = 0; bit < (1 << (N-1)); ++bit) {
    string x;
    x += S.at(0);
    for (int i = 0; i < N-1; ++i) {
      if (bit & (1 << i)) {
        ans += stoll(x);
        x = "";
      }
      x += S.at(i+1);
    }
    ans += stoll(x);
  }
  cout << ans << endl;
}
