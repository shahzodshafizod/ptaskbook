#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix67");
	
	int M, N;
	cin >> M >> N;

	double** matr = new double* [M];
	for (int i = 0; i < M; i++)
	{
		matr[i] = new double [N];
		for (int j = 0; j < N; j++)
			cin >> matr[i][j];
	}

	bool positive;
	for (int j = 0; j < N;)
	{
		positive = true;
		for (int i = 0; i < M; i++)
			if (matr[i][j] < 0)
				positive = false;

		if (positive)
		{
			double** temp = new double* [M];
			int col, row = 0;
			for (int r = 0; r < M; r++)
			{
				temp[row] = new double [N-1];
				col = 0;
				for (int c = 0; c < N; c++)
				{
					if (c == j)
						continue;

					temp[row][col++] = matr[r][c];
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
		else
			j++;
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
