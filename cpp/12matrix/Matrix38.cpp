#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix38");
	
	int M, N;
	cin >> M >> N;

	int** matr = new int* [M];
	
	int rows = 0;
	const int n = 100;
	int cols[n];
	bool ok;
	
	for (int i = 0; i < M; i++)
	{
		matr[i] = new int [N];
		
		memset(cols, 0, n*sizeof(int));
		for (int j = 0; j < N; j++)
		{
			cin >> matr[i][j];
			cols[ matr[i][j] ]++;
		}

		ok = true;
		for (int j = 0; j < N; j++)
			if (cols[ matr[i][j] ] != 1)
				ok = false;
		rows += (int)ok;
	}
	cout << rows;

	for (int i = 0; i < M; i++)
		delete [] matr[i];
	delete [] matr;
	matr = 0;

	return 0;
}
