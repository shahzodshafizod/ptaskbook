#include <iostream>
using namespace std;

int main()
{
	//Task("Array113");
	
	int N;
	cin >> N;

	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	double maximal;
	int index;
	for (int i = N-1; i > 0; i--)
	{
		index = i;
		maximal = mass[i];
		for (int j = i-1; j >= 0; j--)
			if (mass[j] > maximal)
			{
				maximal = mass[j];
				index = j;
			}
		mass[index] = mass[i];
		mass[i] = maximal;

		for (int j = 0; j < N; j++)
			cout << mass[j];
	}

	delete [] mass;
	mass = 0;
	
	return 0;
}
