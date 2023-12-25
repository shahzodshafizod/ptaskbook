#include <stdio.h>

void main() {
	char fileName[50];
	gets(fileName);
	FILE* f = fopen(fileName, "w");
	int j, N;
	scanf("%d", &N);
	char C;
	for (int i = 1; i <= N; ++i) {
		C = 'A';
		for (j = 1; j <= i; ++j) {
			fprintf(f, "%c", C);
			++C;
		}
		while (j <= N) {
			fprintf(f, "%c", '*');
			++j;
		}
		fprintf(f, "\n");
	}
	fclose(f);
}