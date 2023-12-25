#include "TNode.h"
#include "Tree.h"
#include <iostream>
using namespace std;

int main() {
	Tree tree;
	tree.make();
	tree.display();
	TNode* P1 = tree.getNode();
	TNode* PL = P1 == NULL ? NULL : P1->Left;
	TNode* PR = P1 == NULL ? NULL : P1->Right;
	TNode* P0 = P1 == NULL ? NULL : P1->Parent;
	TNode* P2 = NULL;
	if (P0 != NULL) {
		P2 = P1 == P0->Left ? P0->Right : P0->Left;
	}
	cout << PL << endl;
	cout << PR << endl;
	cout << P0 << endl;
	cout << P2 << endl;
	return 0;
}