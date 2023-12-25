#include <iostream>
using namespace std;

int main()
{
	//Task("Array114");
	
	int N;
	cin >> N;

	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	double temp;
	for (int i = 1; i < N; i++)
	{
		for (int j = i; j > 0; j--)
		{
			if (mass[j] < mass[j-1])
			{
				temp = mass[j];
				mass[j] = mass[j-1];
				mass[j-1] = temp;
			}
		}
		for (int j = 0; j < N; j++)
			cout << mass[j];
	}

	delete [] mass;
	mass = 0;
	
	return 0;
}
