#include <stdio.h>

void main() {
	char fileName[50];
	gets(fileName);
	FILE* f = fopen(fileName, "wb");
	double A, D;
	scanf("%lf%lf", &A, &D);
	int i;
	double elem = A;
	for (i = 1; i <= 10; ++i) {
		fwrite(&elem, sizeof(double), 1, f);
		elem += D;
	}
	fclose(f);
}