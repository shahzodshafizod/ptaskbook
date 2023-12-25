#include <iostream>
using namespace std;

void Inv(double A[], int N);

int main()
{
	//Task("Param4");
	int N;
	double* mass = NULL;

	for (int i = 1; i <= 3; i++)
	{
		cin >> N;
		mass = new double [N];
		for (int j = 0; j < N; j++)
			cin >> mass[j];

		Inv(mass, N);

		for (int j = 0; j < N; j++)
			cout << mass[j];

		delete [] mass;
		mass = NULL;
	}
	
	return 0;
}

void Inv(double A[], int N)
{
	double temp;
	for (int i = 0; i < N/2; i++)
	{
		temp = A[i];
		A[i] = A[N-1-i];
		A[N-1-i] = temp;
	}
}