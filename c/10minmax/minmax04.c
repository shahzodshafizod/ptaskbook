#include <stdio.h>

int main() {
	int N;
	scanf("%d", &N);
	double number, minimal;
	int index = 0;
	for (int i = 1; i <= N; ++i) {
		scanf("%lf", &number);
		if (index == 0) {
			minimal = number;
			index = i;
		} else if (number < minimal) {
			minimal = number;
			index = i;
		}
	}
	printf("%d", index);
	return 0;
}