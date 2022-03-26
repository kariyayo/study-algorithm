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
      // 既に同じグループ
      return false;
    }

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

  int ans = 0;
  for (int k = 0; k < M; ++k) {
    UnionFind uf(N+1);
    // k番目の辺以外でunion-findを構築する
    for (int i = 0; i < M; ++i) {
      if (i != k) {
        uf.unite(as[i], bs[i]);
      }
    }
    // k番目の辺の両端が同じグループではない場合、橋である
    if (!uf.issame(as[k], bs[k])) {
      ++ans;
    }
  }
  cout << ans << endl;
}


template <typename T>
void printVector(vector<T> xs) {
  cout << '[';
  for (auto& x: xs) {
    cout << x << ' ';
  }
  cout << ']' << endl;
}
