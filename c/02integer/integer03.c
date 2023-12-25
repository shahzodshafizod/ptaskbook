#include <stdio.h>

void main() {
	int bytes;
	scanf("%d", &bytes);
	int KBytes = bytes / 1024;
	printf("%d", KBytes);
}