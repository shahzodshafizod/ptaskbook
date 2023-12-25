#include <stdio.h>

int main() {
	double price;
	scanf("%lf", &price);
	for (int i = 1; i <= 10; ++i) {
		printf("%.2f\t", price * i);
	}
	return 0;
}