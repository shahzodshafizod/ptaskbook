#include <stdio.h>
#include <stdlib.h>
void MinmaxNum(double[], int, int*, int*);

int main(void) {
	int N;
	double* mass = NULL;
	int maxIndex, minIndex;
	for (int i = 0; i < 3; ++i) {
		scanf("%d", &N);
		mass = malloc(N * sizeof(double));
		for (int j = 0; j < N; ++j) {
			scanf("%lf", mass+j);
		}
		MinmaxNum(mass, N, &minIndex, &maxIndex);
		printf("%d\t%d\n", minIndex, maxIndex);
		free(mass);
		mass = NULL;
	}
	return 0;
}

void MinmaxNum(double A[], int N, int* NMin, int* NMax) {
	*NMin = *NMax = 0;
	for (int i = 1; i < N; ++i) {
		if (A[i] > A[*NMax]) {
			*NMax = i;
		}
		if (A[i] < A[*NMin]) {
			*NMin = i;
		}
	}
	++(*NMin);
	++(*NMax);
}