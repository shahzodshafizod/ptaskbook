#include <stdio.h>

void main() {
	double A, B;
	scanf("%lf%lf", &A, &B);
	int counter = 0;
	while (A >= B) {
		A -= B;
		++counter;
	}
	printf("%d", counter);
}