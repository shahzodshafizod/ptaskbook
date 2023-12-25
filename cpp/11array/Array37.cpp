#include <iostream>
using namespace std;

int main()
{
	//Task("Array37");
	
	int N;
	cin >> N;
	
	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	int count = 0;
	bool progress = false;
	for (int i = 1; i < N; i++)
	{
		if (mass[i] > mass[i-1])
			progress = true;
		else
		{
			if (progress)
			{
				count++;
				progress = false;
			}
		}
	}
	if (progress)
		count++;

	cout << count;

	delete [] mass;
	mass = 0;
	
	return 0;
}
