#include <bits/stdc++.h>
using namespace std;
using uint = unsigned int;
using ll = long long;
using ull = unsigned long long;
template <typename T> inline bool chmin(T& a, T b) { if (a > b) { a = b; return true; } return false; }
template <typename T> inline bool chmax(T& a, T b) { if (a < b) { a = b; return true; } return false; }
template <typename T> void printVector(vector<T> xs, string prefix = "[ ", string suffix = " ]");

const long long INF_LL = 1LL<<60;
const int INF = 1000010000;


vector<vector<int>> G;
vector<ll> counts;
vector<bool> seen;
vector<ll> ans;

int main() {
  int n;
  cin >> n;

  vector<vector<int>> G(n+1);
  for (int i = 1; i <= n; i++) {
    int v, k;
    cin >> v >> k;
    G[v] = vector<int>(k);
    for (int j = 0; j < k; j++) {
      cin >> G[v][j];
    }
  }

  queue<int> todo;
  vector<int> seen(n+1, -1);

  todo.push(1);
  seen[1] = 0;
  while (!todo.empty()) {
    int cur = todo.front();
    todo.pop();
    int d = seen[cur] + 1;
    vector<int> children = G[cur];
    for (int i = 0; i < (int)children.size(); i++) {
      int next = children[i];
      if (seen[next] != -1) {
        continue;
      }
      todo.push(next);
      seen[next] = d;
    }
  }

  for (int i = 1; i <= n; i++) {
    cout << i << " " << seen[i] << endl;
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
