#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix65");
	
	int M, N;
	cin >> M >> N;

	double** matr = new double* [M];
	for (int i = 0; i < M; i++)
	{
		matr[i] = new double [N];
		for (int j = 0; j < N; j++)
			cin >> matr[i][j];
	}

	int colIndex = -1;
	bool positive;
	for (int j = 0; j < N; j++)
	{
		positive = true;
		for (int i = 0; i < M; i++)
			if (matr[i][j] < 0)
				positive = false;

		if (positive)
		{
			colIndex = j;
			break;
		}
	}

	if (colIndex >= 0)
	{
		double** temp = new double* [M];
		int col, row = 0;
		for (int i = 0; i < M; i++)
		{
			temp[row] = new double [N-1];
			col = 0;
			for (int j = 0; j < N; j++)
			{
				if (j == colIndex)
					continue;

				temp[row][col++] = matr[i][j];
			}
			row++;
		}

		for (int i = 0; i < M; i++)
			delete [] matr[i];
		delete [] matr;
		matr = temp;
		temp = 0;
		N--;
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
