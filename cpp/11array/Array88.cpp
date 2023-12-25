#include <iostream>
using namespace std;

int main()
{
	//Task("Array88");
	
	int N;
	cin >> N;
	
	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	int index = N-1;
	double temp;
	while ( (index > 0) && (mass[index] < mass[index-1]) )
	{
		temp = mass[index];
		mass[index] = mass[index-1];
		mass[index-1] = temp;
		index--;
	}

	for (int i = 0; i < N; i++)
		cout << mass[i];

	delete [] mass;
	mass = 0;
	
	return 0;
}
