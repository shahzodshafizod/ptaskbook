#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix43");
	
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
	bool regress;
	for (int j = 0; j < N; j++)
	{
		regress = true;
		for (int i = 1; i < M; i++)
			if (matr[i-1][j] < matr[i][j])
				regress = false;
		count += (int)regress;
	}
	cout << count;

	for (int i = 0; i < M; i++)
		delete [] matr[i];
	delete [] matr;
	matr = 0;

	return 0;
}
