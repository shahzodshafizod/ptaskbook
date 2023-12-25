#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix41");
	
	int M, N;
	cin >> M >> N;

	int** matr = new int* [M];
	for (int i = 0; i < M; i++)
	{
		matr[i] = new int [N];
		for (int j = 0; j < N; j++)
			cin >> matr[i][j];
	}

	const int n = 101;
	int rows[n];
	int colIndex;
	int maximal, colMaximal;
	for (int j = 0; j < N; j++)
	{
		memset(rows, 0, n*sizeof(int));
		for (int i = 0; i < M; i++)
			rows[ matr[i][j] ]++;
		
		maximal = rows[0];
		for (int i = 1; i < n; i++)
			if (rows[i] > maximal)
				maximal = rows[i];

		if (j == 0)
		{
			colIndex = j;
			colMaximal = maximal;
		}
		else if (maximal > colMaximal)
		{
			colIndex = j;
			colMaximal = maximal;
		}
	}
	cout << colIndex+1;

	for (int i = 0; i < M; i++)
		delete [] matr[i];
	delete [] matr;
	matr = 0;
	
	return 0;
}
