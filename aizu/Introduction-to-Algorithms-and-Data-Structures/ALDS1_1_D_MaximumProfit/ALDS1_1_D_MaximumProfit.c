#include <stdio.h>
#include <stdlib.h>
#include <limits.h>

int main(int argc, char *argv[])
{
  char row[12];
  int i;
  int max_diff = INT_MIN;
  int min;
  int n = atoi(fgets(row, sizeof(row), stdin));

  if (n == 0) {
    printf("err\n");
    exit(1);
  }

  min = atoi(fgets(row, sizeof(row), stdin));
  for (i = 1; i < n; i++) {
    int diff;
    int x = atoi(fgets(row, sizeof(row), stdin));
    diff = x - min;
    if (max_diff < diff) {
      max_diff = diff;
    }
    if (x < min) {
      min = x;
    }
  }

  printf("%d\n", max_diff);
  return 0;
}
