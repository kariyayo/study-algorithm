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
  int N, M;
  cin >> N >> M;
  vector<vector<int>> seen(N, vector<int>(N, -1));
  queue<pair<int, int>> todo;
  todo.push(pair(0, 0));
  seen[0][0] = 0;
  int dMax = sqrt(M);
  while (!todo.empty()) {
    pair<int, int> cur = todo.front();
    todo.pop();
    int cnt = seen[cur.first][cur.second];
    for (int i = max(0, cur.first - dMax); i < min(N, cur.first + dMax + 1); i++) {
      // M = (x-i)**2 + (y-j)**2 から、iが決まればjは2通りしかない
      int d = sqrt(M - pow(cur.first - i, 2));
      vector<int> js = { cur.second - d, cur.second + d};
      for (int j : js) {
        if (j < 0 || j >= N) continue;
        if (seen[i][j] >= 0) continue; // 訪問済み
        int d = pow(i-cur.first, 2) + pow(j-cur.second, 2);
        if (d == M) {
          seen[i][j] = cnt + 1;
          todo.push(pair(i, j));
        }
      }
    }
  }
  for (int i = 0; i < N; i++) {
    printVector(seen[i], "", "");
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
  cout << suffix << "\n";
}
