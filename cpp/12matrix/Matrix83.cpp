#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix83");
	
	int M;
	cin >> M;

	double** A = new double* [M];
	for (int i = 0; i < M; i++)
	{
		A[i] = new double [M];
		for (int j = 0; j < M; j++)
			cin >> A[i][j];
	}

	double sum;
	int row = 0, col = 0;
	while (col < M)
	{
		sum = 0;

		for (int i = row, j = col; j >= 0; i++, j--)
			sum += A[i][j];

		cout << sum;
		col++;
	}

	col--;
	row++;

	while (row < M)
	{
		sum = 0;

		for (int i = row, j = col; i < M; i++, j--)
			sum += A[i][j];

		cout << sum;
		row++;
	}

	for (int i = 0; i < M; i++)
		delete [] A[i];
	delete [] A;
	A = 0;
	
	return 0;
}
