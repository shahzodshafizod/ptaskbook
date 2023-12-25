#include <stdio.h>

void main() {
	double n, sum = 0;
	for (int i = 1; i <= 10; ++i) {
		scanf("%lf", &n);
		sum += n;
	}
	printf("%.2f", sum);
}