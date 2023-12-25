#include "Tree.h"

int main() {
	Tree tree;
	tree.make();
	tree.display();
	tree.deleteAllChilds();
	tree.display();
	return 0;
}