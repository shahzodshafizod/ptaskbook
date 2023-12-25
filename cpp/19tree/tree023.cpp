#include "Tree.h"

int main() {
	Tree tree;
	tree.make();
	tree.display();
	int minimal = tree.getMinData();
	TNode* node = tree.getFirstNodePrefix(minimal);
	tree.display(node);
	return 0;
}