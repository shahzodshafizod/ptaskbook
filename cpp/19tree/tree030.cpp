#include "Tree.h"
#include <iostream>
using namespace std;

int main() {
	int N;
	cout << "N = ";
	cin >> N;
	Tree tree;
	tree.make(N);
	tree.display();
	return 0;
}