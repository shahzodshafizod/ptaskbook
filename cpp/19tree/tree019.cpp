#include "Tree.h"
#include <iostream>
using namespace std;

int main() {
	Tree tree;
	tree.make();
	tree.display();
	cout << '\n';
	int maximal = tree.getMaxData();
	int count = tree.getNodeCountK(maximal);
	cout << "maximal = " << maximal << endl;
	cout << "count = " << count << endl;
	return 0;
}