#include <stdio.h>
#include <stdlib.h>

void main() {
	int N;
	scanf("%d", &N);
	int* arr = (int*)malloc(N * sizeof(int));
	int number = 1;
	for (int i = 0; i < N; ++i) {
		arr[i] = number;
		number += 2;
	}
	for (int i = 0; i < N; ++i) {
		printf("%d\t", arr[i]);
	}
	free(arr);
}