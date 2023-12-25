#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix98");
	
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
	int row, col;
	for (int k = 0; k < M/2; k++)
	{
		row = col = k;
		
		while (col < M-k)
		{
			temp = A[row][col];
			A[row][col] = A[M-1-row][M-1-col];
			A[M-1-row][M-1-col] = temp;
			col++;
		}
		
		row++;
		col--;
		
		while (row < M-1-k)
		{
			temp = A[row][col];
			A[row][col] = A[M-1-row][M-1-col];
			A[M-1-row][M-1-col] = temp;
			row++;
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
