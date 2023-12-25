#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix7");
	
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
	
	for (int j = 0; j < N; j++)
		cout << matr[K-1][j];

	for (int i = 0; i < M; i++)
		delete [] matr[i];

	delete [] matr;
	matr = 0;
	
	return 0;
}
