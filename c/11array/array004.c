#include <stdio.h>
#include <stdlib.h>

int main() {
	int N;
	scanf("%d", &N);
	double A, D;
	scanf("%lf%lf", &A, &D);
	double* array = (double*)malloc(N * sizeof(double));
	for (int i = 0; i < N; ++i) {
		if (i == 0) { array[i] = A; }
		else { array[i] = array[i-1] * D; }
	}
	for (int i = 0; i < N; ++i) {
		printf("%.2f\t", array[i]);
	}
	free(array);
	return 0;
}