#include <stdio.h>

void main() {
	int A, B;
	scanf("%d%d", &A, &B);
	int N = 0;
	for (int i = B-1; i > A; --i) {
		printf("%d\t", i);
		++N;
	}
	printf("\n%d", N);
}