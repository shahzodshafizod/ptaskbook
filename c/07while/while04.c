#include <stdio.h>
#include <stdbool.h>

int main() {
	int N;
	scanf("%d", &N);
	int degree = 1;
	while (degree < N) {
		degree *= 3;
	}
	bool result = degree == N;
	printf("%d", result);
	return 0;
}