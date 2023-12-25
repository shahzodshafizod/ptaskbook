#include <stdio.h>

int main() {
	int N;
	scanf("%d", &N);
	double number, sum = 0, zarb = 1;
	for (int i = 1; i <= N; ++i) {
		scanf("%lf", &number);
		sum += number;
		zarb *= number;
	}
	printf("%.2f\t%.2f", sum, zarb);
	return 0;
}