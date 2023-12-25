#include <iostream>
using namespace std;

double Norm1(double** A, int M, int N);

int main()
{
	//Task("Param19");
	int M, N;
	cin >> M >> N;
	double** A = new double* [M];
	for (int i = 0; i < M; i++)
	{
		A[i] = new double [N];
		
		for (int j = 0; j < N; j++)
			cin >> A[i][j];

		cout << Norm1(A, i, N);
	}

	for (int i = 0; i < M; i++)
		delete [] A[i];
	delete [] A;
	A = NULL;
	
	return 0;
}

double Norm1(double** A, int M, int N)
{
	double sum, maxSum = 0;
	for (int j = 0; j < N; j++)
	{
		sum = 0;
		for (int i = 0; i <= M; i++)
			sum += A[i][j];

		if (maxSum == 0)
			maxSum = sum;
		else if (sum > maxSum)
			maxSum = sum;
	}
	return maxSum;
}