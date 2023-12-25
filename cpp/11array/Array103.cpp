#include <iostream>
using namespace std;

int main()
{
	//Task("Array103");
	
	int N;
	cin >> N;

	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	int iMin, iMax;
	double minimal, maximal;
	maximal = minimal = mass[0];
	iMin = iMax = 0;
	for (int i = 1; i < N; i++)
	{
		if (mass[i] < minimal)
		{
			minimal = mass[i];
			iMin = i;
		}

		if (mass[i] > maximal)
		{
			maximal = mass[i];
			iMax = i;
		}
	}

	double* temp = new double [N+2];
	int index = 0;
	for (int i = 0; i < N; i++)
	{
		if (i == iMin)
			temp[index++] = 0;
		
		temp[index++] = mass[i];

		if (i == iMax)
			temp[index++] = 0;
	}

	N += 2;
	delete [] mass;
	mass = temp;
	temp = 0;
	for (int i = 0; i < N; i++)
		cout << mass[i];

	delete [] mass;
	mass = 0;
	
	return 0;
}
