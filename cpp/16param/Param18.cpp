#include <iostream>
using namespace std;

void ChessBoard(int M, int N, int*** A);

int main()
{
	//Task("Param18");
	int M, N;
	cin >> M >> N;
	int** A = NULL;

	ChessBoard(M, N, &A);

	for (int i = 0; i < M; i++)
		for (int j = 0; j < N; j++)
			cout << A[i][j];

	for (int i = 0; i < M; i++)
		delete [] A[i];
	delete [] A;
	A = NULL;
	
	return 0;
}

void ChessBoard(int M, int N, int*** A)
{
	(*A) = new int* [M];
	for (int i = 0; i < M; i++)
	{
		(*A)[i] = new int [N];
		for (int j = 0; j < N; j++)
			(*A)[i][j] = (i+j) % 2 == 0 ? 0 : 1;
	}
}