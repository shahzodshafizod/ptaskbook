#include <stdio.h>

void main() {
	int A, B;
	scanf("%d%d", &A, &B);
	int N = 0;
	for (int i = A; i <= B; ++i) {
		printf("%d\t", i);
		++N;
	}
	printf("\nN = %d", N);
}