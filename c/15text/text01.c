#include <stdio.h>

void main() {
	char fileName[50];
	gets(fileName);
	FILE* f = fopen(fileName, "w");
	int N, K;
	scanf("%d%d", &N, &K);
	for (int i = 1; i <= N; ++i) {
		for (int j = 1; j <= K; ++j) {
			fprintf(f, "%c", '*');
		}
		fprintf(f, "\n");
	}
	fclose(f);
}