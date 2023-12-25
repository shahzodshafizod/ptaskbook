#include <stdio.h>

void main() {
	char S[50];
	gets(S);
	FILE* f = fopen(S, "wb");
	printf("%d", f != NULL);
	if (f != NULL) {
		fclose(f);
	}
}