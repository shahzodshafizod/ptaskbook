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
	bool areRelatives = false;
	for (TNode* i = P1; i != NULL; i = i->Parent) {
		if (i == P2) {
			areRelatives = true;
			break;
		}
	}
	int L1, L2;
	if (!areRelatives) {
		L1 = L2 = 0;
	} else {
		L1 = tree.getNodeHeight(P1) - 1;
		L2 = tree.getNodeHeight(P2) - 1;
	}
	cout << L1 - L2;
	return 0;
}