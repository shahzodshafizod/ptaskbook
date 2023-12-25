#include <stdio.h>
#include <stdbool.h>

void main() {
	int A;
	scanf("%d", &A);
	bool isPositive = A > 0;
	printf("%d", isPositive);
}