#include <iostream>
using namespace std;

int main()
{
	//Task("Array61");
	
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
		for (int j = i; j < N; j++)
			sum += A[j];

		B[i] = sum / (N - i);
	}

	for (int i = 0; i < N; i++)
		cout << B[i];

	delete [] A;
	delete [] B;
	A = B = 0;
	
	return 0;
}
