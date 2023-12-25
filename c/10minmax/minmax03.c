#include <stdio.h>

void main() {
	int N;
	scanf("%d", &N);
	double a, b, P, maxP;
	scanf("%lf%lf", &a, &b);
	maxP = 2 * (a+b);
	int i;
	for (i = 2; i <= N; ++i) {
		scanf("%lf%lf", &a, &b);
		P = 2 * (a+b);
		if (P > maxP) {
			maxP = P;
		}
	}
	printf("%.2f", maxP);
}