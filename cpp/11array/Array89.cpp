#include <iostream>
using namespace std;

int main()
{
	//Task("Array89");
	
	int N;
	cin >> N;

	double* mass = new double [N];
	int index = 0;
	for (int i = 0; i < N; i++)
	{
		cin >> mass[i];
		if (i > 0)
		{
			if (mass[i] > mass[i-1])
				index = i;
		}
	}

	double temp;
	
	// move to left from index
	for (int i = index; i > 0; --i)
	{
		if (mass[i] > mass[i-1])
		{
			temp = mass[i];
			mass[i] = mass[i-1];
			mass[i-1] = temp;
		}
	}

	// move to right from index-1
	for (int i = index-1; i < N-1; i++)
	{
		if (mass[i] < mass[i+1])
		{
			temp = mass[i];
			mass[i] = mass[i+1];
			mass[i+1] = temp;
		}
	}

	for (int i = 0; i < N; i++)
		cout << mass[i];

	delete [] mass;
	mass = 0;
	
	return 0;
}
