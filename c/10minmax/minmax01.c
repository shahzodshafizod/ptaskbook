#include <stdio.h>

void main() {
	int N;
	scanf("%d", &N);
	double n, maximal, minimal;
	scanf("%lf", &n);
	maximal = minimal = n;
	int i = 2;
	for (; i <= N; ++i) {
		scanf("%lf", &n);
		if (n > maximal) maximal = n;
		if (n < minimal) minimal = n;
	}
	printf("%.2f\n%.2f", minimal, maximal);
}