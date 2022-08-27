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
vector<ll> counts;
vector<bool> seen;
vector<ll> ans;

void dfs(int cur, ll adds) {
  vector<int> childs = G[cur];
  for (int i = 0; i < (int)childs.size(); i++) {
    int next = childs[i];
    if (seen[next]) {
      continue;
    }
    seen[next] = true;
    ans[next] += (adds + counts[next]);
    dfs(next, ans[next]);
  }
}

int main() {
  int N, Q;
  cin >> N >> Q;

  G = vector<vector<int>>(N+1);
  for (int i = 0; i < N-1; i++) {
    int a, b;
    cin >> a >> b;
    G[a].push_back(b);
    G[b].push_back(a);
  }

  counts = vector<ll>(N+1);
  for (int i = 0; i < Q; i++) {
    int p, x;
    cin >> p >> x;
    counts[p] += x;
  }

  seen = vector<bool>(N+1);
  ans = vector<ll>(N+1);
  ll adds = 0;
  for (int i = 1; i <= N; i++) {
    if (seen[i]) {
      continue;
    }
    seen[i] = true;
    ans[i] = counts[i] + adds;
    dfs(i, ans[i]);
  }

  for (int i = 1; i <= N; i++) {
    cout << ans[i];
    if (i != N) {
      cout << " ";
    }
  }
  cout << endl;
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
