#include <stdio.h>

int main() {
	const double PI = 3.14;
	double d, L;
	scanf("%lf", &d);
	L = PI * d;
	printf("%.2f", L);
	return 0;
}