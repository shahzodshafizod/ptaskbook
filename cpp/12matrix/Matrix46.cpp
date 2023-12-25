#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix46");
	
	int M, N;
	cin >> M >> N;

	int** matr = new int* [M];
	for (int i = 0; i < M; i++)
	{
		matr[i] = new int [N];
		for (int j = 0; j < N; j++)
			cin >> matr[i][j];
	}

	int maximal, minimal, maxmin = 0;
	int col, row;
	for (int i = 0; i < M; i++)
	{
		maximal = matr[i][0];
		col = 0;
		for (int j = 1; j < N; j++)
		{
			if (matr[i][j] > maximal)
			{
				maximal = matr[i][j];
				col = j;
			}
		}
		minimal = matr[i][col];
		row = i;
		for (int r = 0; r < M; r++)
		{
			if (matr[r][col] < minimal)
			{
				minimal = matr[r][col];
				row = r;
			}
		}
		if (row == i)
		{
			maxmin = matr[row][col];
			break;
		}
	}
	
	cout << maxmin;

	for (int i = 0; i < M; i++)
		delete [] matr[i];
	delete [] matr;
	matr = 0;

	return 0;
}
