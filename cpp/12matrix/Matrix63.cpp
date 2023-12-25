#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix63");
	
	int M, N;
	cin >> M >> N;

	double** matr = new double* [M];
	double minimal;
	int rowIndex = -1;
	for (int i = 0; i < M; i++)
	{
		matr[i] = new double [N];
		
		for (int j = 0; j < N; j++)
		{
			cin >> matr[i][j];

			if (rowIndex < 0)
			{
				minimal = matr[i][j];
				rowIndex = i;
			}
			else if (matr[i][j] < minimal)
			{
				minimal = matr[i][j];
				rowIndex = i;
			}
		}
	}

	if (rowIndex >= 0)
	{
		double** temp = new double* [M-1];
		int col, row = 0;
		for (int i = 0; i < M; i++)
		{
			if (i == rowIndex)
				continue;

			temp[row] = new double [N];
			col = 0;
			for (int j = 0; j < N; j++)
			{
				temp[row][col++] = matr[i][j];
			}
			row++;
		}

		for (int i = 0; i < M; i++)
			delete [] matr[i];
		delete [] matr;
		matr = temp;
		temp = 0;
		M--;
	}

	for (int i = 0; i < M; i++)
		for (int j = 0; j < N; j++)
			cout << matr[i][j];

	for (int i = 0; i < M; i++)
		delete [] matr[i];
	delete [] matr;
	matr = 0;
	
	return 0;
}
