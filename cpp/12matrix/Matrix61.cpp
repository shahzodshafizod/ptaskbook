#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix61");
	
	int M, N;
	cin >> M >> N;

	double** matr = new double* [M];
	for (int i = 0; i < M; i++)
	{
		matr[i] = new double [N];
		for (int j = 0; j < N; j++)
			cin >> matr[i][j];
	}

	int K;
	cin >> K;
	double** temp = new double* [M-1];
	int col, row = 0;
	for (int i = 0; i < M; i++)
	{
		if (i == K-1)
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

	for (int i = 0; i < M; i++)
		for (int j = 0; j < N; j++)
			cout << matr[i][j];

	for (int i = 0; i < M; i++)
		delete [] matr[i];
	delete [] matr;
	matr = 0;
	
	return 0;
}
