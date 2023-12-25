#include "TNode.h"
#include "Tree.h"
#include <iostream>
using namespace std;

int main() {
	Tree tree;
	tree.make();
	tree.display();
	TNode* P1 = tree.getNode();
	TNode* P2 = tree.getRoot(P1);
	cout << P1 << endl;
	cout << P2 << endl;
	return 0;
}