#include <stdio.h>
#include <stdbool.h>
#include "Flag.h"

int main() {
	enum Flag a = DATA;
	printf("%d", sizeof(a));
	return 0;
}