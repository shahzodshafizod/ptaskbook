#include <stdio.h>
#include <math.h>

void Mean(double, double, double*, double*);

void main() {
	double a, b, c, d, am, gm;
	scanf("%lf%lf%lf%lf", &a, &b, &c, &d);
	Mean(a, b, &am, &gm);
	printf("%.2f\t%.2f\n", am, gm);
	Mean(a, c, &am, &gm);
	printf("%.2f\t%.2f\n", am, gm);
	Mean(a, d, &am, &gm);
	printf("%.2f\t%.2f\n", am, gm);
}

void Mean(double X, double Y, double* AMean, double* GMean) {
	*AMean = (X+Y) / 2;
	*GMean = sqrt(X*Y);
}