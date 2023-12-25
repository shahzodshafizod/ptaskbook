#include <stdio.h>

void main() {
	int n;
	scanf("%d", &n);
	n += n > 0 ? -8 : n < 0 ? 6 : 10;
	printf("%d", n);
}