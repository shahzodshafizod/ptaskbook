#include <iostream>
using namespace std;

void RemoveCols(double*** A, int M, int& N, int K1, int K2);

int main()
{
	//Task("Param27");
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
	RemoveCols(&A, M, N, K1, K2);

	cout << M << N;
	for (int i = 0; i < M; i++)
		for (int j = 0; j < N; j++)
			cout << A[i][j];

	for (int i = 0; i < M; i++)
		delete [] A[i];
	delete [] A;
	A = NULL;

	return 0;
}

void RemoveCols(double*** A, int M, int& N, int K1, int K2)
{
	if (K1 > N)
		return;
	if (K2 > N)
		K2 = N;

	int newN = N - (K2 - K1 + 1);
	double** temp = new double* [M];
	for (int i = 0; i < M; i++)
		temp[i] = new double [newN];

	for (int j = 0, k = 0; j < N; j++)
	{
		if (j == K1-1)
		{
			j = K2;
			if (j == N)
				break;
		}
		for (int i = 0; i < M; i++)
			temp[i][k] = (*A)[i][j];
		k++;
	}

	for (int i = 0; i < M; i++)
		delete [] (*A)[i];
	delete [] (*A);
	(*A) = temp;
	temp = NULL;
	N = newN;
}