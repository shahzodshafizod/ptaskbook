#include <iostream>
using namespace std;

double Norm2(double** A, int M, int N, double maxSum);

int main()
{
	//Task("Param20");
	int M, N;
	cin >> M >> N;

	double** A = new double* [M];
	for (int i = 0; i < M; i++)
	{
		A[i] = new double [N];
		for (int j = 0; j < N; j++)
			cin >> A[i][j];
	}

	double maxSum = 0;
	for (int i = 0; i < M; i++)
	{
		maxSum = Norm2(A, i, N, maxSum);
		cout << maxSum;
	}

	for (int i = 0; i < M; i++)
		delete [] A[i];
	delete [] A;
	A = NULL;
	
	return 0;
}

double Norm2(double** A, int M, int N, double maxSum)
{
	double sum = 0;
	for (int j = 0; j < N; j++)
		sum += A[M][j];
	
	return maxSum > sum ? maxSum : sum;
}