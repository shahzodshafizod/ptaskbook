#include <stdio.h>

double PowerN(double, int);

int main() {
	double x;
	int n, i;
	scanf("%lf", &x);
	for (i = 0; i < 5; ++i) {
		scanf("%d", &n);
		printf("%f\n", PowerN(x, n));
	}
	return 0;
}

double PowerN(double X, int N) {
	if (N == 0) {
		return 1;
	}
	if (N < 0) {
		return 1 / PowerN(X, -N);
	}
	if (N % 2 == 0) {
		double half = PowerN(X, N/2);
		return half * half;
	}
	return X * PowerN(X, N-1);
}