#include <iostream>
using namespace std;

void Smooth1(double A[], int N);

int main()
{
	//Task("Param5");
	int N;
	cin >> N;

	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	for (int i = 1; i <= 5; i++)
	{
		Smooth1(mass, N);
		for (int j = 0; j < N; j++)
			cout << mass[j];
	}

	delete [] mass;
	mass = NULL;
	
	return 0;
}

void Smooth1(double A[], int N)
{
	double* temp = new double [N];
	for (int i = 0; i < N; i++)
		temp[i] = A[i];

	double sum;
	for (int i = 0; i < N; i++)
	{
		sum = 0;
		for (int j = 0; j <= i; j++)
			sum += temp[j];
		A[i] = sum / (i+1);
	}

	delete [] temp;
	temp = NULL;
}