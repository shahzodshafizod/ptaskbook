#include <stdio.h>
#include <stdlib.h>

void main() {
	int N;
	scanf("%d", &N);
	double A, D;
	scanf("%lf%lf", &A, &D);
	double* mass = (double*)malloc(N * sizeof(double));
	for (int i = 0; i < N; ++i) {
		mass[i] = A + i*D;
	}
	for (int i = 0; i < N; ++i) {
		printf("%.2f\t", mass[i]);
	}
	free(mass);
}