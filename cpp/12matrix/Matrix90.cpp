#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix90");
	
	int M;
	cin >> M;

	double** matr = new double* [M];
	for (int i = 0; i < M; i++)
	{
		matr[i] = new double [M];
		for (int j = 0; j < M; j++)
			cin >> matr[i][j];
	}

	int row = 0, col = M-1;
	while (row < M)
	{
		for (int i = row, j = col; i < M; i++, j--)
			matr[i][j] = 0;
		
		row++;
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
