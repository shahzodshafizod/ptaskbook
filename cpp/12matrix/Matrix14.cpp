#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix14");
	
	int M;
	cin >> M;

	double** A = new double* [M];
	for (int i = 0; i < M; i++)
	{
		A[i] = new double [M];
		for (int j = 0; j < M; j++)
			cin >> A[i][j];
	}

	for (int j = 0; j < M; j++)
	{
		for (int i = 0; i < M-j; i++)
			cout << A[i][j];

		for (int i = j+1; i < M; i++)
			cout << A[M-1-j][i];
	}

	for (int i = 0; i < M; i++)
		delete [] A[i];
	delete [] A;
	A = 0;
	
	return 0;
}
