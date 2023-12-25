#include <stdio.h>

int Fib1(int);
int count;

int main() {
	int n;
	for (int i = 0; i < 5; ++i) {
		scanf("%d", &n);
		count = 0;
		printf("%d\t", Fib1(n));
		printf("%d\n", count);
	}
	return 0;
}

int Fib1(int N) {
	++count;
	if (N <= 2) {
		return 1;
	}
	return Fib1(N - 1) + Fib1(N - 2);
}