#include "Tree.h"
#include <iostream>
using namespace std;

int main() {
	Tree tree;
	tree.make();
	tree.display();
	cout << '\n';
	int N;
	cout << "N = ";
	cin >> N;
	int index = 0;
	tree.postfixFromN(index, N);
	return 0;
}