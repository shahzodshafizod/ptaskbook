#include <stdio.h>

int main() {
	int N;
	scanf("%d", &N);
	char C = 'A';
	for (int i = 0; i < N; ++i) {
		printf("%c\t", C);
		C++;
	}
	return 0;
}