#include <iostream>
using namespace std;

void DoubleX(int** A, int& N, int X);

int main()
{
	//Task("Param10");
	int X, N, *mass = NULL;

	for (int i = 1; i <= 3; i++)
	{
		cin >> X >> N;
		mass = new int [N];
		for (int j = 0; j < N; j++)
			cin >> mass[j];
		
		DoubleX(&mass, N, X);
		cout << N;
		for (int j = 0; j < N; j++)
			cout << mass[j];

		delete [] mass;
		mass = NULL;
	}
	
	return 0;
}

void DoubleX(int** A, int& N, int X)
{
	for (int i = 0; i < N; i++)
	{
		if ((*A)[i] == X)
		{
			int* temp = new int [N+1];
			for (int k = 0, j = 0; k < N; k++)
			{
				temp[j++] = (*A)[k];
				if (k == i)
					temp[j++] = (*A)[k];
			}
			delete [] (*A);
			(*A) = temp;
			temp = NULL;
			N++;
			i++;
		}
	}
}