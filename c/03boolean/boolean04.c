#include <stdio.h>
#include <stdbool.h>

int main() {
	int A, B;
	scanf("%d%d", &A, &B);
	bool result = A > 2 && B <= 3;
	printf("%d", result);
	return 0;
}