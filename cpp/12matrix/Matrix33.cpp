#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix33");
	
	int M, N;
	cin >> M >> N;

	int** matr = new int* [M];
	for (int i = 0; i < M; i++)
	{
		matr[i] = new int [N];
		for (int j = 0; j < N; j++)
			cin >> matr[i][j];
	}

	int positives, negatives, col = -1;
	for (int j = 0; j < N; j++)
	{
		positives = negatives = 0;
		for (int i = 0; i < M; i++)
			if (matr[i][j] > 0)
				positives++;
			else if (matr[i][j] < 0)
				negatives++;
		if (positives == negatives)
			col = j;
	}

	cout << col+1;

	for (int i = 0; i < M; i++)
		delete [] matr[i];
	delete [] matr;
	matr = 0;

	return 0;
}
