#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix39");
	
	int M, N;
	cin >> M >> N;

	int** matr = new int* [M];
	for (int i = 0; i < M; i++)
	{
		matr[i] = new int [N];
		for (int j = 0; j < N; j++)
			cin >> matr[i][j];
	}

	const int n = 100;
	int rows[n];
	int cols = 0;
	bool ok;
	for (int j = 0; j < N; j++)
	{
		memset(rows, 0, n*sizeof(int));
		for (int i = 0; i < M; i++)
			rows[ matr[i][j] ]++;
		ok = true;
		for (int i = 0; i < M; i++)
			if (rows[ matr[i][j] ] != 1)
				ok = false;
		cols += (int)ok;
	}
	cout << cols;

	for (int i = 0; i < M; i++)
		delete [] matr[i];
	delete [] matr;
	matr = 0;

	return 0;
}
