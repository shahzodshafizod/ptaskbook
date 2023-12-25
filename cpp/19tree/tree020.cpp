#include "Tree.h"
#include <iostream>
using namespace std;

int main() {
	Tree tree;
	tree.make();
	tree.display();
	cout << '\n';
	int minimal = tree.getMinData();
	int count = tree.getLeafCountK(minimal);
	cout << "minimal = " << minimal << endl;
	cout << "leafCount = " << count << endl;
	return 0;
}