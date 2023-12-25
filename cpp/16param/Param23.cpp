#include <iostream>
using namespace std;

void SwapRow(double** A, int M, int N, int K1, int K2);

int main()
{
	//Task("Param23");
	int M, N;
	cin >> M >> N;
	double** A = new double* [M];
	for (int i = 0; i < M; i++)
	{
		A[i] = new double [N];
		for (int j = 0; j < N; j++)
			cin >> A[i][j];
	}

	int K1, K2;
	cin >> K1 >> K2;

	SwapRow(A, M, N, K1, K2);

	for (int i = 0; i < M; i++)
		for (int j = 0; j < N; j++)
			cout << A[i][j];

	for (int i = 0; i < M; i++)
		delete [] A[i];
	delete [] A;
	A = NULL;

	return 0;
}

void SwapRow(double** A, int M, int N, int K1, int K2)
{
	if ( (K1 >= 1) && (K1 <= M) && (K2 >= 1) && (K2 <= M) )
	{
		double temp;
		for (int j = 0; j < N; j++)
		{
			temp = A[K1-1][j];
			A[K1-1][j] = A[K2-1][j];
			A[K2-1][j] = temp;
		}
	}
}