#include <iostream>
using namespace std;

void RemoveRowCol(double*** A, int& M, int& N, int K, int L);

int main()
{
	//Task("Param28");
	int M, N;
	cin >> M >> N;
	double** A = new double* [M];
	for (int i = 0; i < M; i++)
	{
		A[i] = new double [N];
		for (int j = 0; j < N; j++)
			cin >> A[i][j];
	}

	int K, L;
	cin >> K >> L;
	RemoveRowCol(&A, M, N, K, L);

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

void RemoveRowCol(double*** A, int& M, int& N, int K, int L)
{
	if ( (K > M) || (L > N) )
		return;
	
	double** temp = new double* [M-1];
	for (int i = 0; i < M-1; i++)
		temp[i] = new double [N-1];

	for (int i = 0, I = 0; i < M; i++)
	{
		if (i == K-1)
			continue;
		for (int j = 0, J = 0; j < N; j++)
		{
			if (j == L-1)
				continue;
			temp[I][J] = (*A)[i][j];
			J++;
		}
		I++;
	}

	for (int i = 0; i < M; i++)
		delete [] (*A)[i];
	delete [] (*A);
	(*A) = temp;
	temp = NULL;
	M--;
	N--;
}