#include <iostream>
using namespace std;

int main()
{
	//Task("Array53");
	
	int N;
	cin >> N;

	double* A = new double [N];
	double* B = new double [N];
	double* C = new double [N];

	for (int i = 0; i < N; i++)
		cin >> A[i];

	for (int i = 0; i < N; i++)
	{
		cin >> B[i];
		
		C[i] = (A[i] > B[i]) ? A[i] : B[i];
	}

	for (int i = 0; i < N; i++)
		cout << C[i];

	delete [] A;
	delete [] B;
	delete [] C;
	A = B = C = 0;
	
	return 0;
}
