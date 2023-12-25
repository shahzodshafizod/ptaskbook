#include <iostream>
using namespace std;

int main()
{
	//Task("Array116");
	
	int N;
	cin >> N;

	int* A = new int [N];
	cin >> A[0];
	int series = 1;
	int element = A[0];
	for (int i = 1; i < N; i++)
	{
		cin >> A[i];

		if (A[i] != element)
		{
			element = A[i];
			series++;
		}
	}

	int* B = new int [series];
	int* C = new int [series];
	int index = 0;
	B[index] = 1;
	C[index] = A[0];
	for (int i = 1; i < N; i++)
	{
		if (A[i] == C[index])
			B[index]++;
		else
		{
			index++;
			B[index] = 1;
			C[index] = A[i];
		}
	}

	for (int i = 0; i < series; i++)
		cout << B[i];
	for (int i = 0; i < series; i++)
		cout << C[i];

	delete [] A;
	delete [] B;
	delete [] C;
	A = B = C = 0;
	
	return 0;
}
