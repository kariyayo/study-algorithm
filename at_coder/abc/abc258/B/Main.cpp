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


pair<int, int> calc_pos(int N, int x, int y, int dx, int dy) {
  int new_x = x + dx;
  if (new_x > N-1) {
    new_x = new_x - N;
  } else if (new_x < 0) {
    new_x = N + new_x;
  }
  int new_y = y + dy;
  if (new_y > N-1) {
    new_y = new_y - N;
  } else if (new_y < 0) {
    new_y = N + new_y;
  }
  return {new_x, new_y};
}

int main() {
  int N;
  cin >> N;
  vector<vector<int>> as(N);
  for (int i = 0; i < N; i++) {
    as[i] = vector<int>(N);
    for (int j = 0; j < N; j++) {
      char v;
      cin >> v;
      int a = v-'0';
      as[i][j] = a;
    }
  }

  vector<pair<int, int>> move(8);
  move[0] = pair(-1, 0);
  move[1] = pair(-1, 1);
  move[2] = pair(0, 1);
  move[3] = pair(1, 1);
  move[4] = pair(1, 0);
  move[5] = pair(1, -1);
  move[6] = pair(0, -1);
  move[7] = pair(-1, -1);

  ll ans = 0;
  for (int i = 0; i < N; i++) {
    for (int j = 0; j < N; j++) {
      for (int d = 0; d < 8; d++) {
        int x = i;
        int y = j;
        ll v = as[x][y];
        auto [dx, dy] = move[d];
        for (int k = 1; k < N; k++) {
          v = v*10;
          auto [new_x, new_y] = calc_pos(N, x, y, dx, dy);
          x = new_x;
          y = new_y;
          v = v + as[x][y];
        }
        chmax(ans, v);
      }
    }
  }
  cout << ans << endl;
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
