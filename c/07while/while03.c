#include <stdio.h>

void main() {
	int N, K;
	scanf("%d%d", &N, &K);
	int counter = 0;
	while (N >= K) {
		N -= K;
		++counter;
	}
	printf("%d\n", counter);
	printf("%d\n", N);
}