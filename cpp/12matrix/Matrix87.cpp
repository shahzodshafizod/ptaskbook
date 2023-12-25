#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix87");
	
	int M;
	cin >> M;

	double** A = new double* [M];
	for (int i = 0; i < M; i++)
	{
		A[i] = new double [M];
		for (int j = 0; j < M; j++)
			cin >> A[i][j];
	}

	double maximal;
	int row = 0, col = 0;

	while (col < M)
	{
		maximal = A[row][col];

		for (int i = row+1, j = col-1; j >= 0; i++, j--)
			if (A[i][j] > maximal)
				maximal = A[i][j];

		cout << maximal;
		col++;
	}

	row++;
	col--;

	while (row < M)
	{
		maximal = A[row][col];

		for (int i = row+1, j = col-1; i < M; i++, j--)
			if (A[i][j] > maximal)
				maximal = A[i][j];

		cout << maximal;
		row++;
	}

	for (int i = 0; i < M; i++)
		delete [] A[i];
	delete [] A;
	A = 0;
	
	return 0;
}
