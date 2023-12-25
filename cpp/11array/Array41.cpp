#include <iostream>
using namespace std;

int main()
{
	//Task("Array41");
	int N;
	cin >> N;
	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];
	
	double maxSum = mass[0] + mass[1];
	int index = 0;
	for (int i = 1; i < N-1; i++)
	{
		if (mass[i] + mass[i+1] > maxSum)
		{
			maxSum = mass[i] + mass[i+1];
			index = i;
		}
	}

	cout << mass[index] << mass[index+1];

	delete [] mass;
	mass = 0;
	
	return 0;
}
