#include <stdio.h>

void main() {
	char fileName[50];
	gets(fileName);
	FILE* f = fopen(fileName, "w");
	int N;
	scanf("%d", &N);
	char c;
	for (int i = 1; i <= N; ++i) {
		c = 'a';
		for (int j = 1; j <= i; ++j) {
			fprintf(f, "%c", c);
			++c;
		}
		fprintf(f, "\n");
	}
	fclose(f);
}