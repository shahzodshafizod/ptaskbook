#include <iostream>
using namespace std;

bool isLocalMaximal(double** matr, double** temp, int M, int N, int i, int j);

int main()
{
	//Task("Matrix75");
	
	int M, N;
	cin >> M >> N;

	double** matr = new double* [M];
	double** temp = new double* [M];
	for (int i = 0; i < M; i++)
	{
		matr[i] = new double [N];
		temp[i] = new double [N];
		for (int j = 0; j < N; j++)
		{
			cin >> matr[i][j];
			temp[i][j] = matr[i][j];
		}
	}

	for (int i = 0; i < M; i++)
		for (int j = 0; j < N; j++)
			if (isLocalMaximal(matr, temp, M, N, i, j))
				matr[i][j] *= -1;

	for (int i = 0; i < M; i++)
		for (int j = 0; j < N; j++)
			cout << matr[i][j];

	for (int i = 0; i < M; i++)
	{
		delete [] matr[i];
		delete [] temp[i];
	}
	delete [] matr;
	delete [] temp;
	matr = 0;
	temp = 0;
	
	return 0;
}

bool isLocalMaximal(double** matr, double** temp, int M, int N, int i, int j)
{
	bool localMaximal = true;

	// top && left
	if ( (i-1 >= 0) && (j-1 >= 0) )
		localMaximal = localMaximal && (temp[i][j] > temp[i-1][j-1]);

	// top && right
	if ( (i-1 >= 0) && (j+1 < N) )
		localMaximal = localMaximal && (temp[i][j] > temp[i-1][j+1]);
	
	// bottom && left
	if ( (i+1 < M) && (j-1 >= 0) )
		localMaximal = localMaximal && (temp[i][j] > temp[i+1][j-1]);
	
	// bottom && right
	if ( (i+1 < M) && (j+1 < N) )
		localMaximal = localMaximal && (temp[i][j] > temp[i+1][j+1]);
	
	// left
	if (j-1 >= 0)
		localMaximal = localMaximal && (temp[i][j] > temp[i][j-1]);
	
	// right
	if (j+1 < N)
		localMaximal = localMaximal && (temp[i][j] > temp[i][j+1]);
	
	// top
	if (i-1 >= 0)
		localMaximal = localMaximal && (temp[i][j] > temp[i-1][j]);
	
	// bottom
	if (i+1 < M)
		localMaximal = localMaximal && (temp[i][j] > temp[i+1][j]);

	return localMaximal;
}