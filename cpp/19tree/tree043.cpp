#include "Tree.h"

int main() {
	Tree tree;
	tree.make();
	tree.display();
	tree.deleteSubTree(TREE43);
	tree.display();
	return 0;
}