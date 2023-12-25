#include "Tree.h"

int main() {
	Tree tree;
	tree.make();
	tree.display();
	tree.addLeaves(TREE45);
	tree.display();
	return 0;
}