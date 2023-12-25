#include <iostream>
using namespace std;

int main()
{
	//Task("Array54");
	
	int N;
	cin >> N;

	int* A = new int [N];
	int count = 0;
	for (int i = 0; i < N; i++)
	{
		cin >> A[i];
		if (A[i] % 2 == 0)
			count++;
	}
	
	cout << count;
	int* B = new int [count];
	for (int i = 0, j = 0; i < N; i++)
		if (A[i] % 2 == 0)
			B[j++] = A[i];

	for (int i = 0; i < count; i++)
		cout << B[i];

	delete [] A;
	delete [] B;
	A = B = 0;
	
	return 0;
}
