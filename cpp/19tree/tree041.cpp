#include "Tree.h"

int main() {
	Tree tree;
	tree.make();
	tree.display();
	tree.deleteAllLeaves();
	tree.display();
	return 0;
}