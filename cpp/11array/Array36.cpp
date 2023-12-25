#include <iostream>
using namespace std;

int main()
{
	//Task("Array36");
	int N;
	cin >> N;
	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	double maximal = 0;
	bool found = false;

	for (int i = 1; i < N-1; i++)
		if ( (mass[i] > mass[i-1]) && (mass[i] < mass[i+1]) || 
			 (mass[i] < mass[i-1]) && (mass[i] > mass[i+1]) )
			 if (found == false)
			 {
				maximal = mass[i];
				found = true;
			 }
			 else if (mass[i] > maximal)
				 maximal = mass[i];

	cout << maximal;

	delete [] mass;
	mass = 0;
	
	return 0;
}
