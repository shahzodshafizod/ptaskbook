#include <stdio.h>
#include <stdlib.h>

void main() {
	int M, N;
	scanf("%d%d", &M, &N);
	int** matr = malloc(M * sizeof(int*));
	for (int i = 0; i < M; ++i)
		matr[i] = malloc(N * sizeof(int));
	
	for (int i = 0; i < M; ++i)
		for (int j = 0; j < N; ++j)
			matr[i][j] = 10 * (i+1);

	for (int i = 0; i < M; ++i) {
		for (int j = 0; j < N; ++j)
			printf("%d\t", matr[i][j]);
		printf("\n");
	}
	
	for (int i = 0; i < M; ++i)
		free(matr[i]);
	free(matr);
}