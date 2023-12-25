#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix86");
	
	int M;
	cin >> M;

	double** A = new double* [M];
	for (int i = 0; i < M; i++)
	{
		A[i] = new double [M];
		for (int j = 0; j < M; j++)
			cin >> A[i][j];
	}

	double minimal;
	int row = 0, col = M-1;
	while (col >= 0)
	{
		minimal = A[row][col];

		for (int i = row+1, j = col+1; j < M; i++, j++)
			if (A[i][j] < minimal)
				minimal = A[i][j];

		cout << minimal;
		col--;
	}

	row++;
	col++;

	while (row < M)
	{
		minimal = A[row][col];

		for (int i = row+1, j = col+1; i < M; i++, j++)
			if (A[i][j] < minimal)
				minimal = A[i][j];

		cout << minimal;
		row++;
	}

	for (int i = 0; i < M; i++)
		delete [] A[i];
	delete [] A;
	A = 0;
	
	return 0;
}
