#include <iostream>
using namespace std;

int main()
{
	//Task("Array115");
	
	int N;
	cin >> N;

	double* A = new double [N];
	int* I = new int [N];
	for (int i = 0; i < N; i++)
	{
		cin >> A[i];
		I[i] = i;
	}

	int temp;
	for (int i = 0; i < N-1; i++)
	{
		for (int j = 1; j < N-i; j++)
		{
			if (A[I[j]] < A[I[j-1]])
			{
				temp = I[j];
				I[j] = I[j-1];
				I[j-1] = temp;
			}
		}
	}

	for (int i = 0; i < N; i++)
		cout << I[i]+1;

	delete [] A;
	delete [] I;
	A = 0;
	I = 0;
	
	return 0;
}
