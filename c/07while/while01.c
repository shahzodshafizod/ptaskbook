#include <stdio.h>

void main() {
	double A, B;
	scanf("%lf%lf", &A, &B);
	while (A >= B) {
		A -= B;
	}
	printf("%.2f", A);
}