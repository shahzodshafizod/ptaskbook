#include <iostream>
using namespace std;

int main()
{
	//Task("Array74");
	
	int N;
	cin >> N;

	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	double minimal, maximal;
	int iMin, iMax;
	minimal = maximal = mass[0];
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

	if (iMin > iMax)
	{
		int temp = iMin;
		iMin = iMax;
		iMax = temp;
	}

	for (int i = iMin+1; i < iMax; i++)
		mass[i] = 0;

	for (int i = 0; i < N; i++)
		cout << mass[i];

	delete [] mass;
	mass = 0;
	
	return 0;
}