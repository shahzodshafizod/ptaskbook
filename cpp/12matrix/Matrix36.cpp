#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix36");
	
	int M, N;
	cin >> M >> N;

	int** matr = new int* [M];
	for (int i = 0; i < M; i++)
	{
		matr[i] = new int [N];
		for (int j = 0; j < N; j++)
			cin >> matr[i][j];
	}

	int count;
	int rows = 0;
	bool hast;
	for (int i = 1; i < M; i++)
	{
		// first with others
		count = 0;
		for (int j = 0; j < N; j++)
		{
			hast = false;
			for (int k = 0; k < N; k++)
			{
				if (matr[0][j] == matr[i][k])
					hast = true;
			}
			count += (int)hast;
		}

		if (count != N)
			continue;

		// others with first
		count = 0;
		for (int j = 0; j < N; j++)
		{
			hast = false;
			for (int k = 0; k < N; k++)
			{
				if (matr[0][k] == matr[i][j])
					hast = true;
			}
			count += (int)hast;
		}
		
		rows += count == N ? 1 : 0;
	}
	cout << rows;

	for (int i = 0; i < M; i++)
		delete [] matr[i];
	delete [] matr;
	matr = 0;

	return 0;
}
