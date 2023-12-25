#include <stdio.h>

void main() {
	int N;
	scanf("%d", &N);
	double a, b, S, minS;
	scanf("%lf%lf", &a, &b);
	minS = a*b;
	int i;
	for (i = 2; i <= N; ++i) {
		scanf("%lf%lf", &a, &b);
		S = a*b;
		if (S < minS) {
			minS = S;
		}
	}
	printf("%.2f", minS);
}