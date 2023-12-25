#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix18");
	
	int M, N;
	cin >> M >> N;

	double** matr = new double* [M];
	for (int i = 0; i < M; i++)
	{
		matr[i] = new double [N];
		for (int j = 0; j < N; j++)
			cin >> matr[i][j];
	}

	int K;
	cin >> K;

	double sum = 0, pro = 1;
	for (int i = 0; i < M; i++)
	{
		sum += matr[i][K-1];
		pro *= matr[i][K-1];
	}
	cout << sum << pro;

	for (int i = 0; i < M; i++)
		delete [] matr[i];
	delete [] matr;
	matr = 0;

	return 0;
}
