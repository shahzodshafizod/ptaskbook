#include <iostream>
using namespace std;

int main()
{
	//Task("Array55");
	
	int N;
	cin >> N;

	int* A = new int [N];
	for (int i = 0; i < N; i++)
		cin >> A[i];

	int sizeB = N/2 + N%2;
	int* B = new int [sizeB];
	for (int i = 0, j = 0; i < N; i += 2)
		B[j++] = A[i];

	cout << sizeB;
	for (int i = 0; i < sizeB; i++)
		cout << B[i];

	delete [] A;
	delete [] B;
	A = B = 0;
	
	return 0;
}
