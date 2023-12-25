#include <stdio.h>
#include <stdlib.h>

int main() {
	int M, N;
	scanf("%d%d", &M, &N);
	double** matrix = (double**)malloc(M * sizeof(double*));
	for (int i = 0; i < M; ++i) {
		matrix[i] = (double*)malloc(N * sizeof(double));
	}
	double number;
	for (int j = 0; j < N; ++j) {
		scanf("%lf", &(matrix[0][j]));
		for (int i = 1; i < M; ++i) {
			matrix[i][j] = matrix[0][j];
		}
	}
	for (int i = 0; i < M; ++i) {
		for (int j = 0; j < N; ++j) {
			printf("%.2f\t", matrix[i][j]);
		}
		printf("\n");
	}
	for (int i = 0; i < M; ++i) {
		free(matrix+i);
	}
	free(matrix);
	return 0;
}