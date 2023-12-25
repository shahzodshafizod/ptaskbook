#include <iostream>
#include <cmath>
using namespace std;

int main()
{
	//Task("Matrix31");
	
	int M, N;
	cin >> M >> N;

	double** matr = new double* [M];
	double AMean = 0;
	for (int i = 0; i < M; i++)
	{
		matr[i] = new double [N];

		for (int j = 0; j < N; j++)
		{
			cin >> matr[i][j];
			AMean += matr[i][j];
		}
	}
	AMean /= (M*N);
	
	int Index = 0, Jndex = 0;
	double minimal = abs(matr[0][0] - AMean);
	for (int i = 0; i < M; i++)
		for (int j = 0; j < N; j++)
			if (abs(matr[i][j] - AMean) < minimal)
			{
				minimal = abs(matr[i][j] - AMean);
				Index = i;
				Jndex = j;
			}

	cout << Index+1 << Jndex+1;

	for (int i = 0; i < M; i++)
		delete [] matr[i];
	delete [] matr;
	matr = 0;

	return 0;
}
