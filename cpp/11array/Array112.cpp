#include <iostream>
using namespace std;

int main()
{
	//Task("Array112");
	
	int N;
	cin >> N;

	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	double temp;
	for (int i = 0; i < N-1; i++)
	{
		for (int j = 1; j < N-i; j++)
			if (mass[j] < mass[j-1])
			{
				temp = mass[j];
				mass[j] = mass[j-1];
				mass[j-1] = temp;
			}
		for (int i = 0; i < N; i++)
			cout << mass[i];
	}

	delete [] mass;
	mass = 0;
	
	return 0;
}
