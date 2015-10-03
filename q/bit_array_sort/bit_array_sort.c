#include <stdio.h>

int main()
{
  unsigned int bit_array = 0;
  unsigned int a = 1;
  unsigned int result;

  FILE *fr;
  FILE *fw;
  char row[10];

  fr = fopen("out/testdata", "r");
  if (!fr) {
    puts("read error");
    exit(1);
  }

  while(fgets(row, sizeof(row), fr) != NULL) {
    int x = atoi(row);
    bit_array = bit_array | (a << x);
  }

  fclose(fr);

  fw = fopen("out/output_bit_array_sort", "w");
  unsigned i;
  for (i = 0; i < 1000000; i++) {
    if(fprintf(fw , "%07u\n", bit_array & i) < 0) {
      puts("write error");
      exit(1);
    }
  }

  fclose(fw);

  return 0;
}

