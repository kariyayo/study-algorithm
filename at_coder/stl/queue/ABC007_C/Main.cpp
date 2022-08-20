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


int main() {
  int R, C;
  int sy, sx;
  int gy, gx;
  cin >> R >> C;
  cin >> sy >> sx;
  cin >> gy >> gx;

  vector<vector<char>> table(R);
  for (int ri = 0; ri < R; ri++) {
    table[ri] = vector<char>(C);
    for (int ci = 0; ci < C; ci++) {
      cin >> table[ri][ci];
    }
  }

  // 各マスまでの最短移動回数
  vector<vector<int>> dist(R, vector<int>(C, INF));

  queue<pair<int, int>> q;
  q.push(make_pair(sy-1, sx-1));
  dist[sy-1][sx-1] = 0;

  while(!q.empty()) {
    auto pos = q.front();
    q.pop();
    int r = pos.first;
    int c = pos.second;

    int dr[4] = { 1, 0, -1, 0 };
    int dc[4] = { 0, 1, 0, -1 };
    for (int i = 0; i < 4; i++) {
      int nextRow = r + dr[i];
      int nextCol = c + dc[i];
      // 盤外だったら無視
      if (nextRow < 0 || R <= nextRow || nextCol < 0 || C <= nextCol) {
        continue;
      }
      // 通路じゃなかったら無視
      if (table[nextRow][nextCol] == '#') {
        continue;
      }
      // 訪問済みだったら無視
      if (dist[nextRow][nextCol] != INF) {
        continue;
      }
      dist[nextRow][nextCol] = dist[r][c]+1;
      q.push(make_pair(nextRow, nextCol));
    }
  }
  cout << dist[gy-1][gx-1] << endl;
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
