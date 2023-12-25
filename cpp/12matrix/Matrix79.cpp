#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix79");
	
	int M, N;
	cin >> M >> N;

	double** matr = new double* [M];
	for (int i = 0; i < M; i++)
	{
		matr[i] = new double [N];
		for (int j = 0; j < N; j++)
			cin >> matr[i][j];
	}

	double* maximals = new double [N];
	int* indexs = new int [N];
	for (int j = 0; j < N; j++)
	{
		maximals[j] = matr[0][j];
		for (int i = 1; i < M; i++)
			if (matr[i][j] > maximals[j])
				maximals[j] = matr[i][j];

		indexs[j] = j;
	}

	int tempIndex;
	for (int i = 0; i < N-1; i++)
	{
		for (int j = 1; j < N-i; j++)
		{
			if (maximals[ indexs[j-1] ] > maximals[ indexs[j] ])
			{
				tempIndex = indexs[j-1];
				indexs[j-1] = indexs[j];
				indexs[j] = tempIndex;
			}
		}
	}

	double** temp = new double* [M];
	for (int i = 0; i < M; i++)
	{
		temp[i] = new double [N];
		for (int j = 0; j < N; j++)
			temp[i][j] = matr[i][ indexs[j] ];
	}

	for (int i = 0; i < M; i++)
	{
		for (int j = 0; j < N; j++)
			matr[i][j] = temp[i][j];
		
		temp[i] = 0;
	}

	for (int i = 0; i < M; i++)
		delete [] temp[i];
	delete [] temp;
	temp = 0;

	for (int i = 0; i < M; i++)
		for (int j = 0; j < N; j++)
			cout << matr[i][j];

	for (int i = 0; i < M; i++)
		delete [] matr[i];
	delete [] matr;
	matr = 0;
	
	return 0;
}
