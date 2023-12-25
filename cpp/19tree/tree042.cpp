#include "Tree.h"

int main() {
	Tree tree;
	tree.make();
	tree.display();
	tree.deleteSubTree(TREE42);
	tree.display();
	return 0;
}