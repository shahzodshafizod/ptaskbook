#include <stdio.h>

void PowerA3(double, double*);

void main() {
	double a, b;
	int i = 1;
	for (; i < 6; ++i) {
		scanf("%lf", &a);
		PowerA3(a, &b);
		printf("%.2f\n", b);
	}
}

void PowerA3(double A, double* B) {
	*B = A*A*A;
}