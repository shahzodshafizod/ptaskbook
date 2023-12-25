#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix51");
	
	int M, N;
	cin >> M >> N;

	double** matr = new double* [M];
	double minimal, maximal;
	int minRow = -1, maxRow = -1;
	for (int i = 0; i < M; i++)
	{
		matr[i] = new double [N];
		for (int j = 0; j < N; j++)
		{
			cin >> matr[i][j];

			if ( (minRow < 0) || (maxRow < 0) )
			{
				maximal = minimal = matr[i][j];
				maxRow = minRow = 0;
			}
			else if (matr[i][j] < minimal)
			{
				minimal = matr[i][j];
				minRow = i;
			}
			else if (matr[i][j] > maximal)
			{
				maximal = matr[i][j];
				maxRow = i;
			}
		}
	}

	if (minRow != maxRow)
	{
		double temp;
		for (int j = 0; j < N; j++)
		{
			temp = matr[minRow][j];
			matr[minRow][j] = matr[maxRow][j];
			matr[maxRow][j] = temp;
		}
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
