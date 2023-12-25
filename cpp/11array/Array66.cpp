#include <iostream>
using namespace std;

int main()
{
	//Task("Array66");
	
	int N;
	cin >> N;

	int* mass = new int [N];
	int number = 0;
	bool found = false;
	for (int i = 0; i < N; i++)
	{
		cin >> mass[i];
		if ( (found == false) && (mass[i] % 2 == 0) )
		{
			number = mass[i];
			found = true;
		}
	}

	for (int i = 0; i < N; i++)
		if (mass[i] % 2 == 0)
			mass[i] += number;

	for (int i = 0; i < N; i++)
		cout << mass[i];

	delete [] mass;
	mass = 0;
	
	return 0;
}
