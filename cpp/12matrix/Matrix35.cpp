#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix35");
	
	int M, N;
	cin >> M >> N;

	int** matr = new int* [M];
	for (int i = 0; i < M; i++)
	{
		matr[i] = new int [N];
		for (int j = 0; j < N; j++)
			cin >> matr[i][j];
	}

	int col = -1;
	bool unevens;
	for (int j = 0; j < N; j++)
	{
		unevens = true;
		for (int i = 0; i < M; i++)
			if (matr[i][j] % 2 == 0)
				unevens = false;
		if (unevens)
		{
			col = j;
			break;
		}
	}
	cout << col+1;

	for (int i = 0; i < M; i++)
		delete [] matr[i];
	delete [] matr;
	matr = 0;

	return 0;
}
