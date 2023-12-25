#include <iostream>
using namespace std;

int main()
{
	//Task("Array120");
	
	int N;
	cin >> N;

	int* mass = new int [N];
	int element, series = 1;
	cin >> mass[0];
	element = mass[0];
	for (int i = 1; i < N; i++)
	{
		cin >> mass[i];

		if (mass[i] != element)
		{
			element = mass[i];
			series++;
		}
	}

	int* temp = new int [N-series];
	int index = 0;
	element = mass[0];
	for (int i = 1; i < N; i++)	
	{
		if (mass[i] != element)
		{
			element = mass[i];
			continue;
		}

		temp[index++] = mass[i];
	}

	delete [] mass;
	mass = temp;
	temp = 0;
	N -= series;

	for (int i = 0; i < N; ++i)
		cout << mass[i];

	delete [] mass;
	mass = 0;
	
	return 0;
}
