#include <stdio.h>
#include <stdlib.h>

int MinElem(int[], int);

void main() {
	int n;
	int* mass = NULL;
	for (int i = 1; i <= 3; ++i) {
		scanf("%d", &n);
		mass = malloc(n * sizeof(int));
		for (int j = 0; j < n; ++j) {
			scanf("%d", mass+j);
		}
		printf("%d\n", MinElem(mass, n));
		free(mass);
		mass = NULL;
	}
}

int MinElem(int A[], int N) {
	int minimal = A[0];
	for (int i = 1; i < N; ++i) {
		if (A[i] < minimal) minimal = A[i];
	}
	return minimal;
}