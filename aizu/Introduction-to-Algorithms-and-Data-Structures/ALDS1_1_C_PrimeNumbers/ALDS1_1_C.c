#include <stdio.h>
#include <math.h>
#define N_MAX 10000

int isPrime(int x)
{
  int i;
  if (x <= 3) {
    return 1;
  } else if (x % 2 == 0) {
    return 0;
  } else {
    for (i = 3; i <= sqrt(x); i = i + 2) {
      if (x % i == 0) {
        return 0;
      }
    }
    return 1;
  }
}

int main(int argc, char *argv[])
{
  int i, j, n, x;
  int answer = 0;
  scanf("%d", &n);
  for (i = 0; i < n; i++) {
    scanf("%d", &x);
    answer = answer + isPrime(x);
  }
  printf("%d\n", answer);
  return 0;
}
