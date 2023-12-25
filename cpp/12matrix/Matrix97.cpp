#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix97");
	
	int M;
	cin >> M;

	double** A = new double* [M];
	for (int i = 0; i < M; i++)
	{
		A[i] = new double [M];
		for (int j = 0; j < M; j++)
			cin >> A[i][j];
	}

	double temp;
	for (int j = 0; j < M-1; j++)
	{
		for (int i = 0; i < M-1-j; i++)
		{
			temp = A[i][j];
			A[i][j] = A[M-1-j][M-1-i];
			A[M-1-j][M-1-i] = temp;
		}
	}

	for (int i = 0; i < M; i++)
		for (int j = 0; j < M; j++)
			cout << A[i][j];

	for (int i = 0; i < M; i++)
		delete [] A[i];
	delete [] A;
	A = 0;
	
	return 0;
}
