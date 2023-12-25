#include <stdio.h>

void main() {
	double a, b;
	scanf("%lf%lf", &a, &b);
	double P = 2 * (a+b);
	double S = a*b;
	printf("%.2f\n%.2f", S, P);
}