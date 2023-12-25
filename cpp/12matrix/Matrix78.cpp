#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix78");
	
	int M, N;
	cin >> M >> N;

	double** matr = new double* [M];
	double* minimals = new double [M];
	int* indexs = new int [M];
	for (int i = 0; i < M; i++)
	{
		matr[i] = new double [N];
		cin >> matr[i][0];
		minimals[i] = matr[i][0];
		for (int j = 1; j < N; j++)
		{
			cin >> matr[i][j];

			if (matr[i][j] < minimals[i])
				minimals[i] = matr[i][j];
		}
		indexs[i] = i;
	}

	int tempIndex;
	for (int i = 0; i < M-1; i++)
	{
		for (int j = 1; j < M-i; j++)
		{
			if ( minimals[ indexs[j-1] ] < minimals[ indexs[j] ] )
			{
				tempIndex = indexs[j-1];
				indexs[j-1] = indexs[j];
				indexs[j] = tempIndex;
			}
		}
	}

	double** temp = new double* [M];
	for (int i = 0; i < M; i++)
		temp[i] = matr[ indexs[i] ];

	for (int i = 0; i < M; i++)
	{
		matr[i] = temp[i];
		temp[i] = 0;
	}
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
