#include <stdio.h>

void main() {
	int n;
	scanf("%d", &n);
	if (n > 0)
		n -= 8;
	printf("%d", n);
}