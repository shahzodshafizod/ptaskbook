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
	int K;
	printf("K = ");
	scanf("%d", &K);
	printf("%d", Tree_getNodeCount(&tree, DATA, K));
	Tree_destr(&tree);
	return 0;
}