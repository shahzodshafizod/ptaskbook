#include <iostream>
using namespace std;

void ArrayToMatrCol(double A[], int K, int M, int N, double*** B);

int main()
{
	//Task("Param17");
	int K;
	cin >> K;
	double* A = new double [K];
	for (int i = 0; i < K; i++)
		cin >> A[i];

	int M, N;
	cin >> M >> N;
	double** B = NULL;

	ArrayToMatrCol(A, K, M, N, &B);

	for (int i = 0; i < M; i++)
		for (int j = 0; j < N; j++)
			cout << B[i][j];

	delete [] A;
	A = NULL;
	for (int i = 0; i < M; i++)
		delete [] B[i];
	delete [] B;
	B = NULL;
	
	return 0;
}

void ArrayToMatrCol(double A[], int K, int M, int N, double*** B)
{
	(*B) = new double* [M];
	for (int i = 0; i < M; i++)
		(*B)[i] = new double [N];

	int index = 0;
	for (int i = 0; i < N; i++)
		for (int j = 0; j < M; j++)
			(*B)[j][i] = index < K ? A[index++] : 0;
}