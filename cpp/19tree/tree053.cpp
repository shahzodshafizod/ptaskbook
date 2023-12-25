#include "TNode.h"
#include "Tree.h"
#include <iostream>
using namespace std;

int main() {
	Tree tree;
	tree.make();
	tree.display();
	cout << "\nSelect node for pointer P1:\n";
	TNode* P1 = tree.getNode();
	cout << "\nSelect node for pointer P2:\n";
	TNode* P2 = tree.getNode();
	TNode* closestCommonParent = tree.getClosestCommonParent(P1, P1, P2, P2);
	cout << closestCommonParent;
	return 0;
}