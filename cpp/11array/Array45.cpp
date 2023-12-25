#include <iostream>
#include <cmath>
using namespace std;

int main()
{
	//Task("Array45");
	
	int N;
	cin >> N;
	
	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	double minimal = abs(mass[0] - mass[1]);
	int indexI = 0, indexJ = 1;
	
	for (int i = 0; i < N-1; i++)
		
		for (int j = i+1; j < N; j++)
			
			if (abs(mass[i] - mass[j]) < minimal)
			{
				minimal = abs(mass[i] - mass[j]);
				indexI = i;
				indexJ = j;
			}

	if (indexI < indexJ)
		cout << indexI+1 << indexJ+1;
	else
		cout << indexJ+1 << indexI+1;

	delete [] mass;
	mass = 0;
	
	return 0;
}
