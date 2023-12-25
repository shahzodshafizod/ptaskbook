#include <iostream>
using namespace std;

void Smooth3(double A[], int N);

int main()
{
	//Task("Param7");
	int N;
	cin >> N;
	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	for (int i = 1; i <= 5; i++)
	{
		Smooth3(mass, N);
		for (int j = 0; j < N; j++)
			cout << mass[j];
	}

	delete [] mass;
	mass = NULL;
	
	return 0;
}

void Smooth3(double A[], int N)
{
	if (N <= 1)
		return;
	
	double* temp = new double [N];
	for (int i = 0; i < N; i++)
		temp[i] = A[i];

	A[0] = (temp[0] + temp[1]) / 2;
	if (N <= 2)
	{
		A[1] = A[0];
		return;
	}
	
	for (int i = 1; i < N-1; i++)
		A[i] = (temp[i-1] + temp[i] + temp[i+1]) / 3;
	A[N-1] = (temp[N-2] + temp[N-1]) / 2;

	delete [] temp;
	temp = NULL;
}