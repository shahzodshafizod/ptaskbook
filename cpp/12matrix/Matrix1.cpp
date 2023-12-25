#include <iostream>
using namespace std;

int main()
{
	//Task("Matrix1");
	
	int M, N;
	cin >> M >> N;

	int** matr = new int* [M];
	for (int i = 0; i < M; i++)
	{
		matr[i] = new int [N];
		
		for (int j = 0; j < N; j++)
			matr[i][j] = 10*(i+1);
	}

	for (int i = 0; i < M; i++)
		for (int j = 0; j < N; j++)
			cout << matr[i][j];

	for (int i = 0; i < M; i++)
		delete [] matr[i];
	delete [] matr;
	matr = 0;
	
	return 0;
}
