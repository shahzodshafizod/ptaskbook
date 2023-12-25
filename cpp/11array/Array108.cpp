#include <iostream>
using namespace std;

int main()
{
	//Task("Array108");
	
	int N;
	cin >> N;

	double* mass = new double [N];
	int positives = 0;
	for (int i = 0; i < N; i++)
	{
		cin >> mass[i];
		
		if (mass[i] > 0)
			positives++;
	}

	double* temp = new double [N+positives];
	int index = 0;
	for (int i = 0; i < N; i++)
	{
		if (mass[i] > 0)
			temp[index++] = 0;

		temp[index++] = mass[i];
	}

	N += positives;
	delete [] mass;
	mass = temp;
	temp = 0;
	
	for (int i = 0; i < N; i++)
		cout << mass[i];

	delete [] mass;
	mass = 0;
	
	return 0;
}
