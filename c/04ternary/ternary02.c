#include <stdio.h>

void main() {
	int n;
	scanf("%d", &n);
	n += n > 0 ? -8 : 6;
	printf("%d", n);
}