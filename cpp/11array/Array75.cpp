#include <iostream>
using namespace std;

int main()
{
	//Task("Array75");
	
	int N;
	cin >> N;

	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	double minimal, maximal;
	minimal = maximal = mass[0];
	int iMin, iMax;
	iMin = iMax = 0;
	for (int i = 0; i < N; i++)
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

	int from = iMin < iMax ? iMin : iMax;
	int to = iMin > iMax ? iMin : iMax;
	int howMany = to - from + 1;
	double temp;
	for (int i = from, j = 0; j < howMany/2; i++, j++)
	{
		temp = mass[i];
		mass[i] = mass[to-j];
		mass[to-j] = temp;
	}

	for (int i = 0; i < N; i++)
		cout << mass[i];

	delete [] mass;
	mass = 0;
	
	return 0;
}
