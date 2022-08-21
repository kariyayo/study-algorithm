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
vector<int> seen;
vector<int> complete;
int t;

void dfs(int node) {
  auto childs = G[node];
  for (int i = 0; i < (int)childs.size(); i++) {
    auto next = childs[i];
    if (seen[next] != -1) {
      continue;
    }
    t++;
    seen[next] = t;
    dfs(next);
    t++;
    complete[next] = t;
  }
}

int main() {
  int n;
  cin >> n;

  G = vector<vector<int>>(n+1);

  for (int i = 0; i < n; i++) {
    int u, k;
    cin >> u >> k;
    G[u] = vector<int>(k);
    for (int j = 0; j < k; j++) {
      cin >> G[u][j];
    }
  }

  seen = vector<int>(n+1, -1);
  complete = vector<int>(n+1, -1);
  t = 0;
  for (int i = 1; i < n+1; i++) {
    if (seen[i] != -1) {
      continue;
    }
    t++;
    seen[i] = t;
    dfs(i);
    t++;
    complete[i] = t;
  }

  for (int id = 1; id <= n; id++) {
    cout << id << " " << seen[id] << " " << complete[id] << endl;
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
