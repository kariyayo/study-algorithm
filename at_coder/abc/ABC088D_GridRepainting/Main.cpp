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
  int H, W;
  cin >> H >> W;

  vector<vector<int>> dist(H);
  int cntBlack = 0;
  vector<vector<char>> table(H);
  for (int row = 0; row < H; row++) {
    table[row] = vector<char>(W);
    dist[row] = vector<int>(W, -1);
    for (int col = 0; col < W; col++) {
      cin >> table[row][col];
      if (table[row][col] == '#') {
        cntBlack++;
      }
    }
  }
  queue<pair<int, int>> todo;
  todo.push(pair(0, 0));
  dist[0][0] = 0;
  while (!todo.empty()) {
    pair<int, int> cur = todo.front();
    todo.pop();
    int d = dist[cur.first][cur.second];
    vector<int> dr = {1, 0, -1, 0};
    vector<int> dc = {0, 1, 0, -1};
    for (int i = 0; i < 4; i++) {
      int nextR = cur.first + dr[i];
      int nextC = cur.second + dc[i];
      if (nextR < 0 || H <= nextR) continue;
      if (nextC < 0 || W <= nextC) continue;
      if (dist[nextR][nextC] != -1) continue;
      if (table[nextR][nextC] == '#') continue;
      todo.push(pair(nextR, nextC));
      dist[nextR][nextC] = d+1;
    }
  }
  if (dist[H-1][W-1] == -1) {
    cout << -1 << endl;
  } else {
    cout << (H*W) - 2 - cntBlack - (dist[H-1][W-1] - 1) << endl;
  }
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
