#include <stdio.h>

void main() {
	int n;
	scanf("%d", &n);
	if (n > 0) {
		n -= 8;
	} else if (n < 0) {
		n += 6;
	} else {
		n = 10;
	}
	printf("%d", n);
}