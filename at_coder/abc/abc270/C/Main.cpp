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
    complete[next] = node;
  }
}

int main() {
  int N, X, Y;
  cin >> N >> X >> Y;

  G = vector<vector<int>>(N+1);
  for (int i = 0; i < N-1; i++) {
    int U, V;
    cin >> U >> V;
    G[U].push_back(V);
    G[V].push_back(U);
  }

  seen = vector<int>(N+1, -1);
  complete = vector<int>(N+1, -1);
  dfs(X);
  vector<int> ans;
  int node = Y;
  while (node != X) {
    ans.push_back(node);
    node = complete[node];
  }
  ans.push_back(X);
  reverse(ans.begin(), ans.end());
  printVector(ans, "", "");
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
