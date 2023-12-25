#include "TNode.h"
#include "Tree.h"
#include <iostream>
using namespace std;

int main() {
	Tree tree1;
	tree1.make();
	tree1.display();
	cout << "\nSelect node for pointer P1:\n";
	TNode* P1 = tree1.getNode();
	P1 = tree1.getRoot(P1);
	
	Tree tree2;
	tree2.copy(P1);
	tree2.display();
	return 0;
}