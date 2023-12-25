#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix32");
	
	int M, N;
	cin >> M >> N;
	
	int** matr = new int* [M];
	int positives, negatives, row = -1;
	for (int i = 0; i < M; i++)
	{
		matr[i] = new int [N];
		
		positives = negatives = 0;
		for (int j = 0; j < N; j++)
		{
			cin >> matr[i][j];
			
			if (matr[i][j] > 0)
				positives++;
			
			else if (matr[i][j] < 0)
				negatives++;
		}

		if ((row < 0) && (positives == negatives))
			row = i;
	}

	cout << row+1;

	for (int i = 0; i < M; i++)
		delete [] matr[i];
	delete [] matr;
	matr = 0;

	return 0;
}
