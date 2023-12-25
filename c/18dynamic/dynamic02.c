#include <stdio.h>
#include "TStack.h"

int main() {
	struct TStack stack;
	TStack_constr(&stack);
	TStack_handmake(&stack);
	int N = 0;
	struct TNode* lastNode = NULL;
	for (struct TNode* i = TStack_getTop(&stack); i != NULL; i = i->Next) {
		lastNode = i;
		printf("%d\t", i->Data);
		++N;
	}
	printf("\nN = %d", N);
	printf("\nlastNode = %p", lastNode);
	TStack_destr(&stack);
	return 0;
}