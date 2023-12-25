#include <iostream>
using namespace std;

int main()
{
	//Task("Array62");
	
	int N;
	cin >> N;

	double* A = new double [N];
	int sizeB = 0, sizeC = 0;
	for (int i = 0; i < N; i++)
	{
		cin >> A[i];

		if (A[i] > 0)
			sizeB++;
		else
			sizeC++;
	}

	double* B = new double [sizeB];
	double* C = new double [sizeC];
	for (int indexA = 0, indexB = 0, indexC = 0; indexA < N; indexA++)
		if (A[indexA] > 0)
			B[indexB++] = A[indexA];
		else
			C[indexC++] = A[indexA];

	cout << sizeB;
	for (int i = 0; i < sizeB; i++)
		cout << B[i];

	cout << sizeC;
	for (int i = 0; i < sizeC; i++)
		cout << C[i];

	delete [] A;
	delete [] B;
	delete [] C;
	A = B = C;
	
	return 0;
}
