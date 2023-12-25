#include <stdio.h>
#include <stdlib.h>

int MaxNum(double[], int);

void main() {
	int n;
	double* mass = NULL;
	for (int i = 1; i <= 3; ++i) {
		scanf("%d", &n);
		mass = malloc(n * sizeof(double));
		for (int j = 0; j < n; ++j) {
			scanf("%lf", mass+j);
		}
		printf("%d\n", MaxNum(mass, n));
		free(mass);
		mass = NULL;
	}
}

int MaxNum(double A[], int N) {
	int index = 0;
	for (int i = 1; i < N; ++i) {
		if (A[i] > A[index]) index = i;
	}
	return index+1;
}