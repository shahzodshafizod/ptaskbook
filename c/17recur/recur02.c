#include <stdio.h>

double Fact2(int);

void main() {
	int n;
	for (int i = 1; i <= 5; ++i) {
		scanf("%d", &n);
		printf("%.1f\n", Fact2(n));
	}
}

double Fact2(int N) {
	if (N <= 1) {
		return 1;
	}
	return N * Fact2(N-2);
}