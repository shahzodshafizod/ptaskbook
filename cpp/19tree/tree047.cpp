#include "Tree.h"

int main() {
	Tree tree;
	tree.make();
	tree.display();
	int L = tree.getLevel();
	tree.doPerfect(0, L);
	tree.display();
	return 0;
}