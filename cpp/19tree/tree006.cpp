#include "Tree.h"
#include <iostream>
using namespace std;

int main() {
	Tree tree;
	tree.make();
	tree.display();
	cout << '\n' << tree.getLeafCount();
	return 0;
}