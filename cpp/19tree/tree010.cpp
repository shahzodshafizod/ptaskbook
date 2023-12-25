#include "Tree.h"
#include <iostream>
using namespace std;

int main() {
	Tree tree;
	tree.make();
	tree.display();
	int level = tree.getLevel();
	int* arr = new int [level+1];
	for (int i = 0; i <= level; ++i) {
		arr[i] = 0;
	}
	tree.levelNodeCountToArr(arr, 0);
	for (int i = 0; i <= level; ++i) {
		cout << arr[i] << '\t';
	}
	delete [] arr;
	arr = NULL;
	return 0;
}