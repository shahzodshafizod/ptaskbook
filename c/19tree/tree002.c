#include <stdio.h>
#include "Tree.h"
#include <stdlib.h>
#include <time.h>

int main() {
	srand(time(0));
	struct Tree tree;
	Tree_constr(&tree);
	Tree_automake(&tree, rand() % 7 + 7);
	Tree_display(&tree);
	printf("\n%d", Tree_getNodeCount(&tree, ALL, 0));
	Tree_destr(&tree);
	return 0;
}