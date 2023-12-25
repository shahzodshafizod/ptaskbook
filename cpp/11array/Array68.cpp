#include <iostream>
using namespace std;

int main()
{
	//Task("Array68");
	
	int N;
	cin >> N;

	double* mass = new double[N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	double minimal, maximal;
	int iMin, iMax;
	minimal = maximal = mass[0];
	iMin = iMax = 0;
	for (int i = 1; i < N; i++)
	{
		if (mass[i] > maximal)
		{
			maximal = mass[i];
			iMax = i;
		}

		if (mass[i] < minimal)
		{
			minimal = mass[i];
			iMin = i;
		}
	}
	double temp = mass[iMin];
	mass[iMin] = mass[iMax];
	mass[iMax] = temp;

	for (int i = 0; i < N; i++)
		cout << mass[i];

	delete [] mass;
	mass = 0;
	
	return 0;
}
