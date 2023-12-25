#include <iostream>
using namespace std;

void Transp(double** A, int M);

int main()
{
	//Task("Param25");
	int M;
	cin >> M;
	double** A = new double* [M];
	for (int i = 0; i < M; i++)
	{
		A[i] = new double [M];
		for (int j = 0; j < M; j++)
			cin >> A[i][j];
	}

	Transp(A, M);

	for (int i = 0; i < M; i++)
		for (int j = 0; j < M; j++)
			cout << A[i][j];

	for (int i = 0; i < M; i++)
		delete [] A[i];
	delete [] A;
	A = NULL;

	return 0;
}

void Transp(double** A, int M)
{
	double temp;
	for (int i = 0; i < M; i++)
		for (int j = i + 1; j < M; j++)
		{
			temp = A[i][j];
			A[i][j] = A[j][i];
			A[j][i] = temp;
		}
}