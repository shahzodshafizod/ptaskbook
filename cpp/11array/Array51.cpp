#include <iostream>
using namespace std;

int main()
{
	//Task("Array51");
	
	int N;
	cin >> N;
	
	double* A = new double [N];
	double* B = new double [N];

	for (int i = 0; i < N; i++)
		cin >> A[i];
	
	for (int i = 0; i < N; i++)
		cin >> B[i];

	double temp;
	for (int i = 0; i < N; i++)
	{
		temp = A[i];
		A[i] = B[i];
		B[i] = temp;
	}

	for (int i = 0; i < N; i++)
		cout << A[i];

	for (int i = 0; i < N; i++)
		cout << B[i];

	delete [] A;
	delete [] B;
	A = B = 0;
	
	return 0;
}
