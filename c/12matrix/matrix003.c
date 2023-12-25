#include <stdio.h>
#include <stdlib.h>

void main() {
	int M, N;
	scanf("%d%d", &M, &N);
	double** matr = malloc(M * sizeof(double*));
	for (int i = 0; i < M; ++i) {
		matr[i] = malloc(N * sizeof(double));
	}
	for (int i = 0; i < M; ++i) {
		scanf("%lf", *(matr+i));
		for (int j = 1; j < N; ++j) {
			matr[i][j] = matr[i][0];
		}
	}
	for (int i = 0; i < M; ++i) {
		for (int j = 0; j < N; ++j) {
			printf("%.2f\t", matr[i][j]);
		}
		printf("\n");
	}
	for (int i = 0; i < M; ++i) {
		free(matr[i]);
	}
	free(matr);
}