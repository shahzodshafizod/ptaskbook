#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix99");
	
	int M;
	cin >> M;

	double** A = new double* [M];
	for (int i = 0; i < M; i++)
	{
		A[i] = new double [M];
		for (int j = 0; j < M; j++)
			cin >> A[i][j];
	}

	double temp;
	int col;
	for (int row = 0; row < M/2; row++)
	{
		col = row;
		while (col < M-1-row)
		{
			temp = A[row][col];
			A[row][col] = A[col][M-1-row];
			A[col][M-1-row] = A[M-1-row][M-1-col];
			A[M-1-row][M-1-col] = A[M-1-col][row];
			A[M-1-col][row] = temp;
			
			col++;
		}
	}

	for (int i = 0; i < M; i++)
		for (int j = 0; j < M; j++)
			cout << A[i][j];

	for (int i = 0; i < M; i++)
		delete [] A[i];
	delete [] A;
	A = 0;
	
	return 0;
}
