#include <stdio.h>

void main() {
	char fileName[50];
	gets(fileName);
	FILE* f = fopen(fileName, "wb");
	int N;
	scanf("%d", &N);
	int i, n = 2;
	for (i = 1; i <= N; ++i) {
		fwrite(&n, sizeof(int), 1, f);
		printf("%d\t", n);
		n += 2;
	}
	fclose(f);
}