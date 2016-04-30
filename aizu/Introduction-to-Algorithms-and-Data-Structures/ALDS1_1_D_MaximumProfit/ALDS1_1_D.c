#include <stdio.h>
#include <limits.h>

int main(int argc, char *argv[])
{
  int i, n, x, min, diff;
  int max_diff = INT_MIN;

  scanf("%d", &n);
  scanf("%d", &min);

  for (i = 1; i < n; i++) {
    scanf("%d", &x);
    diff = x - min;
    if (max_diff < diff) {
      max_diff = diff;
    }
    if (min > x) {
      min = x;
    }
  }

  printf("%d\n", max_diff);

  return 0;
}
