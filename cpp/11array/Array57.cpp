#include <iostream>
using namespace std;

int main()
{
	//Task("Array57");
	
	int N;
	cin >> N;

	int* A = new int [N];
	int* B = new int [N];
	for (int i = 0; i < N; i++)
		cin >> A[i];

	int index = 0;
	for (int i = 1; i < N; i += 2)
		B[index++] = A[i];

	for (int i = 0; i < N; i += 2)
		B[index++] = A[i];

	for (int i = 0; i < N; i++)
		cout << B[i];

	delete [] A;
	delete [] B;
	A = B = 0;
	
	return 0;
}
