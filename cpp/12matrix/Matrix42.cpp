#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix42");
	
	int M, N;
	cin >> M >> N;

	double** matr = new double* [M];
	for (int i = 0; i < M; i++)
	{
		matr[i] = new double [N];
		for (int j = 0; j < N; j++)
			cin >> matr[i][j];
	}

	int count = 0;
	bool progress;
	for (int i = 0; i < M; i++)
	{
		progress = true;
		
		for (int j = 1; j < N; j++)
			if (matr[i][j-1] > matr[i][j])
				progress = false;
		
		count += (int)progress;
	}
	cout << count;

	for (int i = 0; i < M; i++)
		delete [] matr[i];
	delete [] matr;
	matr = 0;

	return 0;
}
