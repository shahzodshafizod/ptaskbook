#include <stdio.h>

void PowerA234(double, double*, double*, double*);

void main() {
	double a, b, c, d;
	int i;
	for (i = 1; i <= 5; ++i) {
		scanf("%lf", &a);
		PowerA234(a, &b, &c, &d);
		printf("%.2f\t%.2f\t%.2f\n", b, c, d);
	}
}

void PowerA234(double A, double* B, double* C, double* D) {
	*B = A*A;
	*C = (*B) * A;
	*D = (*C) * A;
}