#include <stdio.h>
#include "Tree.h"
#include <stdlib.h>
#include <string.h>
#include <time.h>

int main() {
	srand(time(0));
	struct Tree tree;
	Tree_constr(&tree);
	Tree_automake(&tree, rand() % 7 + 7);
	Tree_display(&tree);
	int size = Tree_getLevel(&tree) + 1;
	int* array = (int*)malloc(size * sizeof(int));
	memset(array, 0, size * sizeof(int));
	printf("\n");
	Tree_toArray(&tree, array, 0, SUM);
	for (int i = 0; i < size; ++i) {
		printf("%d\t", *(array+i));
	}
	Tree_destr(&tree);
	return 0;
}