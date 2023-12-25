#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix37");
	
	int M, N;
	cin >> M >> N;

	int** matr = new int* [M];
	for (int i = 0; i < M; i++)
	{
		matr[i] = new int [N];
		for (int j = 0; j < N; j++)
			cin >> matr[i][j];
	}

	int count, cols = 0;
	bool hast;
	for (int j = 0; j < N-1; j++)
	{
		// last with others
		count = 0;
		for (int i = 0; i < M; i++)
		{
			hast = false;
			for (int k = 0; k < M; k++)
				if (matr[i][N-1] == matr[k][j])
					hast = true;
			count += (int)hast;
		}
		if (count != M)
			continue;

		// others with last
		count = 0;
		for (int i = 0; i < M; i++)
		{
			hast = false;
			for (int k = 0; k < M; k++)
				if (matr[i][j] == matr[k][N-1])
					hast = true;
			count += (int)hast;
		}

		cols += count == M ? 1 : 0;
	}
	cout << cols;

	for (int i = 0; i < M; i++)
		delete [] matr[i];
	delete [] matr;
	matr = 0;

	return 0;
}
