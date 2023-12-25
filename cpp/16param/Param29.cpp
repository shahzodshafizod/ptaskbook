#include <iostream>
using namespace std;

void SortCols(int** A, int M, int N);
void SwapCols(int** A, int M, int col1, int col2);

int main()
{
	//Task("Param29");
	int M, N;
	cin >> M >> N;
	int** A = new int* [M];
	for (int i = 0; i < M; i++)
	{
		A[i] = new int [N];
		for (int j = 0; j < N; j++)
			cin >> A[i][j];
	}

	SortCols(A, M, N);

	for (int i = 0; i < M; i++)
		for (int j = 0; j < N; j++)
			cout << A[i][j];

	for (int i = 0; i < M; i++)
		delete [] A[i];
	delete [] A;
	A = NULL;

	return 0;
}

void SortCols(int** A, int M, int N)
{
	bool areEqual;
	for (int j = 0; j < N-1; j++)
		for (int k = 1; k < N-j; k++)
		{
			areEqual = true;
			for (int i = 0; i < M; i++)
			{
				if (A[i][k-1] < A[i][k])
					break;
				if (A[i][k-1] != A[i][k])
					areEqual = false;
			}
			if (!areEqual)
				SwapCols(A,M, k-1, k);
		}
}

void SwapCols(int** A, int M, int col1, int col2)
{
	int temp;
	for (int i = 0; i < M; i++)
	{
		temp = A[i][col1];
		A[i][col1] = A[i][col2];
		A[i][col2] = temp;
	}
}