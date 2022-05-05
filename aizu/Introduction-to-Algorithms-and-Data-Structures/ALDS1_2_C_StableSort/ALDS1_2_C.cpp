#include <bits/stdc++.h>
#include <atcoder/all>
using namespace std;
using uint = unsigned int;
using ll = long long;
using ull = unsigned long long;
template <typename T> inline bool chmin(T& a, T b) { if (a > b) { a = b; return true; } return false; }
template <typename T> inline bool chmax(T& a, T b) { if (a < b) { a = b; return true; } return false; }
template <typename T> void printVector(vector<T> xs, string prefix = "[ ", string suffix = " ]");

// const long long INF = 1LL<<60;
const int INF = 1000010000;

bool gt(string v1, string v2) {
  int x1 = v1[1] - '0';
  int x2 = v2[1] - '0';
  if (x1 > x2) {
    return true;
  }
  return false;
}

void bubbleSort(int n, vector<string> &xs) {
  for (int i = 0; i < n; i++) {
    for (int j = n - 1; j > i; j--) {
      if (gt(xs[j-1], xs[j])) {
        string tmp = xs[j-1];
        xs[j-1] = xs[j];
        xs[j] = tmp;
      }
    }
  }
}

void selectionSort(int n, vector<string> &xs) {
  for (int i = 0; i < n; i++) {
    int min_idx = i;
    for (int j = i; j < n; j++) {
      if (gt(xs[min_idx], xs[j])) {
        min_idx = j;
      }
    }
    if (min_idx != i) {
      string tmp = xs[i];
      xs[i] = xs[min_idx];
      xs[min_idx] = tmp;
    }
  }
}

int main() {
  int N;
  cin >> N;
  vector<string> as(N);
  for (int i = 0; i < N; i++) cin >> as[i];
  vector<string> bs(N);
  copy(as.begin(), as.end(), bs.begin());
  vector<string> cs(N);
  copy(as.begin(), as.end(), cs.begin());

  bubbleSort(N, bs);
  printVector(bs, "", "");
  cout << "Stable" << endl;

  selectionSort(N, cs);
  printVector(cs, "", "");
  if (equal(bs.begin(), bs.end(), cs.begin())) {
    cout << "Stable" << endl;
  } else {
    cout << "Not stable" << endl;
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
