#include <iostream>
using namespace std;

void Smooth2(double A[], int N);

int main()
{
	//Task("Param6");
	int N;
	cin >> N;

	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	for (int i = 0; i < 5; i++)
	{
		Smooth2(mass, N);
		for (int j = 0; j < N; j++)
			cout << mass[j];
	}

	delete [] mass;
	mass = NULL;

	return 0;
}

void Smooth2(double A[], int N)
{
	double* temp = new double [N];
	for (int i = 0; i < N; i++)
		temp[i] = A[i];

	for (int i = 1; i < N; i++)
		A[i] = (temp[i-1] + temp[i]) / 2;

	delete [] temp;
	temp = NULL;
}