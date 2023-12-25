#include "Tree.h"
#include <iostream>
using namespace std;

int main() {
	Tree tree;
	tree.make();
	tree.display();
	cout << '\n';
	int level = tree.getLevel();
	int L;
	cout << "L = ";
	cin >> L;
	int N = 0;
	if (L <= level+1) {
		N = tree.getLevelNodeCount(L);
	}
	cout << "\nN = " << N;
	return 0;
}