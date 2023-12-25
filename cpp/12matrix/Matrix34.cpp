#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix34");
	
	int M, N;
	cin >> M >> N;

	int** matr = new int* [M];
	int row = -1;
	bool evens;
	for (int i = 0; i < M; i++)
	{
		matr[i] = new int [N];
		evens = true;
		for (int j = 0; j < N; j++)
		{
			cin >> matr[i][j];
			if (matr[i][j] % 2 != 0)
				evens = false;
		}
		if (evens == true)
			row = i;
	}
	cout << row+1;

	for (int i = 0; i < M; i++)
		delete [] matr[i];
	delete [] matr;
	matr = 0;

	return 0;
}
