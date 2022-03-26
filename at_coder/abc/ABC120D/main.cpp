#include <bits/stdc++.h>
#include <atcoder/all>
using namespace std;
using uint = unsigned int;
using ll = long long;
using ull = unsigned long long;
template <typename T> inline bool chmin(T& a, T b) { if (a > b) { a = b; return true; } return false; }
template <typename T> inline bool chmax(T& a, T b) { if (a < b) { a = b; return true; } return false; }
template <typename T> void printVector(vector<T> xs);

// const long long INF = 1LL<<60;
const int INF = 1000010000;


struct UnionFind {
  vector<int> par, siz;

  UnionFind(int n) : par(n, -1), siz(n, 1) { }

  int root(int x) {
    if (par[x] == -1) return x;
    else return par[x] = root(par[x]); // 経路圧縮する
  }

  bool issame(int x, int y) {
    return root(x) == root(y);
  }

  bool unite(int x, int y) {
    int x_root = root(x);
    int y_root = root(y);

    if (x_root == y_root) {
      return false;
    }

    // union by size
    if (siz[x_root] < siz[y_root]) {
      swap(x_root, y_root);
    }

    par[y_root] = x_root;
    siz[x_root] += siz[y_root];
    return true;
  }

  int size(int x) {
    return siz[root(x)];
  }
};

int main() {
  int N, M;
  cin >> N >> M;

  vector<int> as(M), bs(M);
  for (int i = 0; i < M; ++i) {
    cin >> as[i] >> bs[i];
  }

  vector<ll> ans(M);
  ans[M-1] = (1ll * (N-1) * ((N-1)+1)) / 2;
  UnionFind uf(N+1);

  // 逆に考えて順番に橋をつなげていき不便さを減らしていく
  for (int i = M-1; i >= 1; --i) {
    int a = as[i], b = bs[i];
    if (uf.issame(a, b)) {
      ans[i-1] = ans[i];
    } else {
      ans[i-1] = ans[i] - (uf.size(a) * uf.size(b));
      uf.unite(a, b);
    }
  }
  for (int i = 0; i < M; i++) {
    cout << ans[i] << endl;
  }
}


template <typename T>
void printVector(vector<T> xs) {
  cout << '[';
  for (auto& x: xs) {
    cout << x << ' ';
  }
  cout << ']' << endl;
}
