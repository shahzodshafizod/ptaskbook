#include <stdio.h>

int main() {
	int a, b, c;
	scanf("%d%d%d", &a, &b, &c);
	int count = 0;
	count += a > 0 ? 1 : 0;
	count += b > 0 ? 1 : 0;
	count += c > 0 ? 1 : 0;
	printf("%d", count);
	return 0;
}