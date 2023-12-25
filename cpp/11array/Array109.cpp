#include <iostream>
using namespace std;

int main()
{
	//Task("Array109");
	
	int N;
	cin >> N;

	double* mass = new double [N];
	int negatives = 0;
	for (int i = 0; i < N; i++)
	{
		cin >> mass[i];

		if (mass[i] < 0)
			negatives++;
	}

	double* temp = new double [N+negatives];
	int index = 0;
	for (int i = 0; i < N; i++)
	{
		temp[index++] = mass[i];

		if (mass[i] < 0)
			temp[index++] = 0;
	}

	delete [] mass;
	mass = temp;
	temp = 0;
	N += negatives;

	for (int i = 0; i < N; i++)
		cout << mass[i];

	delete [] mass;
	mass = 0;
	
	return 0;
}
