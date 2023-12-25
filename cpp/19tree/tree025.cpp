#include "Tree.h"
#include <iostream>
using namespace std;

int main() {
	int N;
	cout << "N = ";
	cin >> N;
	int* arr = new int [N];
	for (int i = 0; i < N; ++i) {
		cin >> arr[i];
	}
	Tree tree;
	tree.make(arr, 0, N, TREE25);
	tree.display();
	delete [] arr;
	arr = NULL;
	return 0;
}