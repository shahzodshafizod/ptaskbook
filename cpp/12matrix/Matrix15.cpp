#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix15");
	
	int M;
	cin >> M;

	double** A = new double* [M];
	for (int i = 0; i < M; i++)
	{
		A[i] = new double [M];
		for (int j = 0; j < M; j++)
			cin >> A[i][j];
	}

	int i = 0;
	for (; i < M/2; i++)
	{
		for (int j = i; j < M-i; j++)
			cout << A[i][j];

		for (int j = i+1; j < M-i; j++)
			cout << A[j][M-1-i];

		for (int j = M-1-i-1; j >= i; j--)
			cout << A[M-1-i][j];

		for (int j = M-1-i-1; j > i; j--)
			cout << A[j][i];
	}

	cout << A[i][i];

	for (int i = 0; i < M; i++)
		delete [] A[i];
	delete [] A;
	A = 0;

	return 0;
}
