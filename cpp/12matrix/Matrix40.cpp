#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix40");
	
	int M, N;
	cin >> M >> N;

	int** matr = new int* [M];
	const int n = 1000;
	int cols[n];
	int maximal, rowMaximal;
	int rowIndex = 0;
	for (int i = 0; i < M; i++)
	{
		matr[i] = new int [N];
		
		memset(cols, 0, n*sizeof(int));
		for (int j = 0; j < N; j++)
		{
			cin >> matr[i][j];
			cols[ matr[i][j] ]++;
		}
		
		maximal = cols[0];
		for (int j = 1; j < n; j++)
			if (cols[j] > maximal)
				maximal = cols[j];
		
		if (i == 0)
		{
			rowIndex = 0;
			rowMaximal = maximal;
		}
		else if (maximal >= rowMaximal)
		{
			rowIndex = i;
			rowMaximal = maximal;
		}
	}
	cout << rowIndex+1;

	for (int i = 0; i < M; i++)
		delete [] matr[i];
	delete [] matr;
	matr = 0;

	return 0;
}
