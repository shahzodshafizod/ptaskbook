#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix54");
	
	int M, N;
	cin >> M >> N;

	double** matr = new double* [M];
	for (int i = 0; i < M; i++)
	{
		matr[i] = new double [N];
		for (int j = 0; j < N; j++)
			cin >> matr[i][j];
	}

	int col = N-1;
	bool negative;
	for (int j = 0; j < N-1; j++)
	{
		negative = true;
		
		for (int i = 0; i < M; i++)
			if (matr[i][j] > 0)
				negative = false;
		
		if (negative)
		{
			col = j;
			break;
		}
	}

	if (col != N-1)
	{
		double temp;
		for (int i = 0; i < M; i++)
		{
			temp = matr[i][col];
			matr[i][col] = matr[i][N-1];
			matr[i][N-1] = temp;
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
