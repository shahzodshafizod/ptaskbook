#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix71");
	
	int M, N;
	cin >> M >> N;

	double** matr = new double* [M];
	int colIndex = -1;
	double minimal;
	for (int i = 0; i < M; i++)
	{
		matr[i] = new double [N];
		for (int j = 0; j < N; j++)
		{
			cin >> matr[i][j];

			if (colIndex < 0)
			{
				minimal = matr[i][j];
				colIndex = j;
			}
			else if (matr[i][j] < minimal)
			{
				minimal = matr[i][j];
				colIndex = j;
			}
		}
	}

	double** temp = new double* [M];
	int col, row = 0;
	for (int i = 0; i < M; i++)
	{
		temp[row] = new double [N+1];
		col = 0;
		for (int j = 0; j < N; j++)
		{
			temp[row][col++] = matr[i][j];

			if (j == colIndex)
				temp[row][col++] = matr[i][j];
		}
		row++;
	}
	for (int i = 0; i < M; i++)
		delete [] matr[i];
	delete [] matr;
	matr = temp;
	temp = 0;
	N++;

	for (int i = 0; i < M; i++)
		for (int j = 0; j < N; j++)
			cout << matr[i][j];

	for (int i = 0; i < M; i++)
		delete [] matr[i];
	delete [] matr;
	matr = 0;
	
	return 0;
}
