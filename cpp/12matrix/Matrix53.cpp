#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix53");
	
	int M, N;
	cin >> M >> N;

	double** matr = new double* [M];
	for (int i = 0; i < M; i++)
	{
		matr[i] = new double [N];
		for (int j = 0; j < N; j++)
			cin >> matr[i][j];
	}

	int col = 0;
	bool positive;
	for (int j = 1; j < N; j++)
	{
		positive = true;
		
		for (int i = 0; i < M; i++)
			if (matr[i][j] < 0)
				positive = false;
		
		if (positive)
			col = j;
	}

	if (col != 0)
	{
		double temp;
		for (int i = 0; i < M; i++)
		{
			temp = matr[i][0];
			matr[i][0] = matr[i][col];
			matr[i][col] = temp;
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
