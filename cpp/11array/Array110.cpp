#include <iostream>
using namespace std;

int main()
{
	//Task("Array110");
	
	int N;
	cin >> N;

	int* mass = new int [N];
	int evens = 0;
	for (int i = 0; i < N; i++)
	{
		cin >> mass[i];

		if (mass[i] % 2 == 0)
			evens++;
	}

	int* temp = new int [N + evens];
	int index = 0;
	for (int i = 0; i < N; i++)
	{
		temp[index++] = mass[i];

		if (mass[i] % 2 == 0)
			temp[index++] = mass[i];
	}

	delete [] mass;
	mass = temp;
	temp = 0;
	N += evens;

	for (int i = 0; i < N; i++)
		cout << mass[i];

	delete [] mass;
	mass = 0;
	
	return 0;
}
