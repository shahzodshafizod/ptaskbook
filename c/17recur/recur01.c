#include <stdio.h>

double Fact(int);

void main() {
	int n;
	for (int i = 1; i <= 5; ++i) {
		scanf("%d", &n);
		printf("%.1f\n", Fact(n));
	}
}

double Fact(int N) {
	if (N <= 1) {
		return 1;
	}
	return N * Fact(N-1);
}