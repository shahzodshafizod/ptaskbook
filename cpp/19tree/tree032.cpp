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
	int index = 0;
	tree.make(arr, index, N, N);
	tree.display();
	return 0;
}