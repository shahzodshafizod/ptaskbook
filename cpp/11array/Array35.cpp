#include <iostream>
using namespace std;

int main()
{
	//Task("Array35");
	int N;
	cin >> N;
	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];
	
	double minimal = 0;
	bool found = false;

	if (mass[0] > mass[1])
	{
		minimal = mass[0];
		found = true;
	}

	for (int i = 1; i < N-1; i++)
	{
		if ( (mass[i] > mass[i-1]) && (mass[i] > mass[i+1]) )
			if (found == false)
			{
				minimal = mass[i];
				found = true;
			}
			else if (mass[i] < minimal)
				minimal = mass[i];
	}

	if (mass[N-1] > mass[N-2])
		if (found == false)
		{
			minimal = mass[N-1];
			found = true;
		}
		else if (mass[N-1] < minimal)
			minimal = mass[N-1];

	if (found)
		cout << minimal;
	else
		cout << 0;

	delete [] mass;
	mass = 0;
	
	return 0;
}
