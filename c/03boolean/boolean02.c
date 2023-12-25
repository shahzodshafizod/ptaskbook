#include <stdio.h>
#include <stdbool.h>

void main() {
	int A;
	scanf("%d", &A);
	bool isOdd = A % 2 != 0;
	printf("%d", isOdd);
}