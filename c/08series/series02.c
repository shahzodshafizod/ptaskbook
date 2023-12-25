#include <stdio.h>

void main() {
	double n, mul = 1;
	for (int i = 1; i <= 10; ++i) {
		scanf("%lf", &n);
		mul *= n;
	}
	printf("%.2f", mul);
}