#include <stdio.h>
#include <math.h>

void TrianglePS(double, double*, double*);

int main() {
	double a, p, s;
	for (int i = 1; i < 4; ++i) {
		scanf("%lf", &a);
		TrianglePS(a, &p, &s);
		printf("%.2f\t%.2f\n", p, s);
	}
	return 0;
}

void TrianglePS(double a, double* P, double* S) {
	*P = 3 * a;
	*S = a*a * sqrt(3) / 4;
}