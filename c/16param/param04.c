#include <stdio.h>
#include <stdlib.h>

void Inv(double[], int);

int main() {
	int n;
	double* array = NULL;
	for (int i = 0; i < 3; ++i) {
		scanf("%d", &n);
		array = (double*)malloc(n * sizeof(double));
		for (int j = 0; j < n; ++j) {
			scanf("%lf", array+j);
		}
		Inv(array, n);
		for (int j = 0; j < n; ++j) {
			printf("%.2f\t", *(array+j));
		}
		printf("\n");
		free(array);
	}
	return 0;
}

void Inv(double A[], int N) {
	double temp;
	for (int i = 0, j = N-1; i < j; ++i, --j) {
		temp = A[i];
		A[i] = A[j];
		A[j] = temp;
	}
}