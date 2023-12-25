#include <iostream>
using namespace std;

int main()
{
	//Task("Array34");
	
	int N;
	cin >> N;
	
	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	double maximal = 0;
	bool found = false;

	if (mass[0] < mass[1])
	{
		maximal = mass[0];
		found = true;
	}
	
	for (int i = 1; i < N-1; i++)
		if ( (mass[i] < mass[i-1]) && (mass[i] < mass[i+1]) )
		{
			if (found == false)
			{
				maximal = mass[i];
				found = true;
			}
			else if (mass[i] > maximal)
				maximal = mass[i];
		}

	if ( (mass[N-1] < mass[N-2]) && (mass[N-1] > maximal) )
	{
		if (found == false)
		{
			maximal = mass[N-1];
			found = true;
		}
		else if (mass[N-1] > maximal)
			maximal = mass[N-1];
	}

	if (found)
		cout << maximal;
	else
		cout << 0;
	
	delete [] mass;
	mass = 0;
	
	return 0;
}
