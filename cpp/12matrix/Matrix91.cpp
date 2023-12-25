#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix91");
	
	int M;
	cin >> M;

	double** matr = new double* [M];
	for (int i = 0; i < M; i++)
	{
		matr[i] = new double [M];
		for (int j = 0; j < M; j++)
			cin >> matr[i][j];
	}

	int row = 0, col = 0;
	while (col < M)
	{
		for (int i = row, j = col; j < M; i++, j++)
			matr[i][j] = 0;

		col++;
	}

	for (int i = 0; i < M; i++)
		for (int j = 0; j < M; j++)
			cout << matr[i][j];

	for (int i = 0; i < M; i++)
		delete [] matr[i];
	delete [] matr;
	matr = 0;
	
	return 0;
}
