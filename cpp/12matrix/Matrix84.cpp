#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix84");
	
	int M;
	cin >> M;

	double** A = new double* [M];
	for (int i = 0; i < M; i++)
	{
		A[i] = new double [M];
		for (int j = 0; j < M; j++)
			cin >> A[i][j];
	}

	double AMean;
	int count, row = 0, col = M-1;
	
	while (col >= 0)
	{
		AMean = 0;
		count = 0;

		for (int i = row, j = col; j < M; i++, j++)
		{
			AMean += A[i][j];
			count++;
		}

		AMean /= count;
		cout << AMean;
		col--;
	}

	row++;
	col++;

	while (row < M)
	{
		AMean = 0;
		count = 0;

		for (int i = row, j = col; i < M; i++, j++)
		{
			AMean += A[i][j];
			count++;
		}

		AMean /= count;
		cout << AMean;
		row++;
	}

	for (int i = 0; i < M; i++)
		delete [] A[i];
	delete [] A;
	A = 0;
	
	return 0;
}
