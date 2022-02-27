#include <bits/stdc++.h>
#include <atcoder/all>

using namespace std;

using uint = unsigned int;
using ll = long long;
using ull = unsigned long long;
const int INF = 20002000;

#define rep(i, n) for (int i = 0; i < (int)(n); ++i)
#define rep2(i, start, n) for (int i = (start); i < (int)(n); ++i)

int N;
ll ans;
void count753(ll x, int used) {
  if (x > N) return;
  if (used == 0b111) ++ans;

  count753(x * 10 + 7, used | 0b001);
  count753(x * 10 + 5, used | 0b010);
  count753(x * 10 + 3, used | 0b100);
}

int main() {
  cin >> N;
  count753(0, 0b000);
  cout << ans << endl;
}
