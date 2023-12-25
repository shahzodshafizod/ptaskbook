#include "Tree.h"
#include <iostream>
using namespace std;

int main() {
	Tree tree;
	tree.make();
	tree.display();
	cout << '\n';
	TNode* node = NULL;
	if (tree.hasOddData()) {
		int maximalOdd = tree.getMaxOddData();
		node = tree.getLastNodeInfix(maximalOdd);
		tree.display(node);
	} else {
		cout << node;
	}
	return 0;
}