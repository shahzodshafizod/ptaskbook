#include <iostream>
using namespace std;

int main()
{
	//Task("Array52");
	
	int N;
	cin >> N;
	
	double* A = new double [N];
	double* B = new double [N];
	for (int i = 0; i < N; i++)
	{
		cin >> A[i];

		if (A[i] < 5)
			B[i] = 2 * A[i];
		else
			B[i] = A[i] / 2;
	}

	for (int i = 0; i < N; i++)
		cout << B[i];

	delete [] A;
	delete [] B;
	A = B = 0;
	
	return 0;
}
