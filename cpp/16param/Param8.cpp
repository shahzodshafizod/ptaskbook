#include <iostream>
using namespace std;

void Remove(int** A, int& N, int X);

int main()
{
	//Task("Param8");
	int X, N, *mass = NULL;

	for (int i = 1; i <= 3; i++)
	{
		cin >> X >> N;
		mass = new int [N];
		for (int j = 0; j < N; j++)
			cin >> mass[j];
		
		Remove(&mass, N, X);
		cout << N;
		for (int j = 0; j < N; j++)
			cout << mass[j];
		
		delete [] mass;
		mass = NULL;
	}
	
	return 0;
}

void Remove(int** A, int& N, int X)
{
	int* temp = new int [N];
	int j = 0;
	for (int i = 0; i < N; i++)
		if ((*A)[i] != X)
			temp[j++] = (*A)[i];

	if (j != N)
	{
		delete [] (*A);
		(*A) = new int [j];
		for (int i = 0; i < j; i++)
			(*A)[i] = temp[i];
		N = j;
	}
	delete [] temp;
	temp = NULL;
}