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

vector<string> memo;

string f(int n) {
  if (n == 1) return "1";
  else {
    if (memo[n] == "") {
      memo[n] = f(n-1) + " " + to_string(n) + " " + f(n-1);
    }
    return memo[n];
  }
}

int main() {
  int N;
  cin >> N;

  memo = vector<string>(N+1);
  cout << f(N) << endl;
}


template <typename T>
void printVector(vector<T> xs) {
  cout << '[';
  for (auto& x: xs) {
    cout << x << ' ';
  }
  cout << ']' << endl;
}
