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
	cout << "\nSelect node for pointer P3:\n";
	TNode* P3 = tree.getNode();
	cout << "\nLevel of node P1 is " << tree.getNodeHeight(P1) - 1;
	cout << "\nLevel of node P2 is " << tree.getNodeHeight(P2) - 1;
	cout << "\nLevel of node P3 is " << tree.getNodeHeight(P3) - 1;
	return 0;
}