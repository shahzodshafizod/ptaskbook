#include <iostream>
using namespace std;

void ArrayToMatrRow(double A[], int K, int M, int N, double*** B);

int main()
{
	//Task("Param16");
	int K;
	cin >> K;
	double* A = new double [K];
	for (int i = 0; i < K; i++)
		cin >> A[i];

	int M, N;
	cin >> M >> N;
	double** B = NULL;
	ArrayToMatrRow(A, K, M, N, &B);
	
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

void ArrayToMatrRow(double A[], int K, int M, int N, double*** B)
{
	(*B) = new double* [M];
	int index = 0;
	for (int i = 0; i < M; i++)
	{
		(*B)[i] = new double [N];
		for (int j = 0; j < N; j++)
			(*B)[i][j] = index < K ? A[index++] : 0;
	}
}