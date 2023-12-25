#include <stdio.h>
#include "TStack.h"

int main() {
	struct TStack stack;
	TStack_constr(&stack);
	int N, data;
	scanf("%d", &N);
	for (int i = 0; i < N; ++i) {
		scanf("%d", &data);
		TStack_Push(&stack, data);
	}
	TStack_print(&stack);
	TStack_destr(&stack);
	return 0;
}