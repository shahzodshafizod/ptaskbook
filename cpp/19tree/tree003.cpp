#include "Tree.h"
#include <iostream>
using namespace std;

int main() {
	Tree tree;
	tree.make();
	tree.display();
	int K;
	cout << "\nK = ";
	cin >> K;
	cout << tree.getNodeCountK(K);
	return 0;
}