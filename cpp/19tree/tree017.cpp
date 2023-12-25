#include "Tree.h"
#include <iostream>
using namespace std;

int main() {
	Tree tree;
	tree.make();
	tree.display();
	cout << '\n';
	int N1, N2;
	cout << "N1 = ";
	cin >> N1;
	cout << "N2 = ";
	cin >> N2;
	int index = 0;
	tree.prefixBetween(index, N1, N2);
	return 0;
}