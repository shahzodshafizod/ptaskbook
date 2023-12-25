#include "Tree.h"
#include <iostream>
using namespace std;

int main() {
	int L, N;
	cout << "L = ";
	cin >> L;
	cout << "N = ";
	cin >> N;
	int* arr = new int [N];
	for (int i = 0; i < N; ++i) {
		cin >> arr[i];
	}
	Tree tree;
	int index = 0;
	tree.make(arr, index, N, 0, L);
	tree.display();
	return 0;
}