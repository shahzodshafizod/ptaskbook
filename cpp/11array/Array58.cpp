#include <iostream>
using namespace std;

int main()
{
	//Task("Array58");
	
	int N;
	cin >> N;

	double* A = new double [N];
	double* B = new double [N];
	for (int i = 0; i < N; i++)
		cin >> A[i];

	double sum;
	for (int i = 0; i < N; i++)
	{
		sum = 0;
		for (int j = 0; j <= i; j++)
			sum += A[j];

		B[i] = sum;
	}

	for (int i = 0; i < N; i++)
		cout << B[i];

	delete [] A;
	delete [] B;
	A = B = 0;
	
	return 0;
}
