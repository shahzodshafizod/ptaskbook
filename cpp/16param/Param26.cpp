#include <iostream>
using namespace std;

void RemoveRows(double*** A, int& M, int N, int K1, int K2);

int main()
{
	//Task("Param26");
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
	RemoveRows(&A, M, N, K1, K2);

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

void RemoveRows(double*** A, int& M, int N, int K1, int K2)
{
	if (K1 > M)
		return;
	if (K2 > M)
		K2 = M;

	int newM = M - (K2 - K1 + 1);
	double** temp = new double* [newM];
	for (int i = 0; i < newM; i++)
		temp[i] = new double [N];

	for (int i = 0, k = 0; i < M; i++)
	{
		if (i == K1-1)
		{
			i = K2;
			if (i == M)
				break;
		}
		for (int j = 0; j < N; j++)
			temp[k][j] = (*A)[i][j];
		k++;
	}

	for (int i = 0; i < M; i++)
		delete [] (*A)[i];
	delete [] (*A);
	(*A) = temp;
	temp = NULL;
	M = newM;
}