#include <stdio.h>
#include "TStack.h"

int main() {
	int D;
	scanf("%d", &D);
	struct TStack stack;
	TStack_constr(&stack);
	TStack_automake(&stack, 7);
	TStack_print(&stack);
	TStack_Push(&stack, D);
	TStack_print(&stack);
	TStack_destr(&stack);
	return 0;
}