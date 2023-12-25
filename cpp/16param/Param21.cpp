#include <iostream>
using namespace std;

double SumRow(double** A, int M, int N, int K);

int main()
{
	//Task("Param21");
	int M, N;
	cin >> M >> N;
	double** A = new double* [M];
	for (int i = 0; i < M; i++)
	{
		A[i] = new double [N];
		for (int j = 0; j < N; j++)
			cin >> A[i][j];
	}

	int K;
	for (int i = 1; i <= 3; i++)
	{
		cin >> K;
		cout << SumRow(A, M, N, K);
	}

	for (int i = 0; i < M; i++)
		delete [] A[i];
	delete [] A;
	A = NULL;
	
	return 0;
}

double SumRow(double** A, int M, int N, int K)
{
	double sum = 0;
	if ( (K >= 1) && (K <= M) )
		for (int j = 0; j < N; j++)
			sum += A[K-1][j];

	return sum;
}