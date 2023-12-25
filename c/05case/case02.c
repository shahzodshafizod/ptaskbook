#include <stdio.h>

void main() {
	int K;
	scanf("%d", &K);
	switch (K) {
		case 1: printf("bad"); break;
		case 2: printf("unsatisfactory"); break;
		case 3: printf("mediocre"); break;
		case 4: printf("good"); break;
		case 5: printf("excellent"); break;
		default: printf("error"); break;
	}
}