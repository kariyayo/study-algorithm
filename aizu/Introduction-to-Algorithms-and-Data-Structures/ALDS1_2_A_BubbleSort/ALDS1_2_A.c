#include <stdio.h>

void printArray(int *as, int n)
{
  int i;
  for (i = 0; i < n - 1; i++) {
    printf("%d ", as[i]);
  }
  printf("%d\n", as[n - 1]);
}

void bubbleSort(int *as, int n)
{
  int i, j;
  int x, y;
  int count = 0;

  for (i = 0; i < n; i++) {
    for (j = n - 1; j > i; j--) {
      x = as[j];
      y = as[j - 1];
      if (x < y) {
        as[j] = y;
        as[j - 1] = x;
        count++;
      }
    }
  }
  printArray(as, n);
  printf("%d\n", count);
}

int main(int argc, char *argv[])
{
  char row[4];
  int n = atoi(fgets(row, sizeof(row), stdin));
  int as[n];
  int i;

  for (i = 0; i < n; i++) {
    scanf("%d", &as[i]);
  }

  bubbleSort(as, n);

  return 0;
}
