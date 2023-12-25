#include <iostream>
#include <cmath>
using namespace std;

int main()
{
	//Task("Array46");
	
	double R;
	cin >> R;
	
	int N;
	cin >> N;
	
	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	double closeSum = abs(mass[0] + mass[1] - R);
	int indexI = 0, indexJ = 1;
	
	for (int i = 0; i < N-1; i++)
		
		for (int j = i+1; j < N; j++)
			
			if (abs(mass[i] + mass[j] - R) < closeSum)
			{
				closeSum = abs(mass[i] + mass[j] - R);
				indexI = i;
				indexJ = j;
			}
	
	cout << mass[indexI] << mass[indexJ];
	
	delete [] mass;
	mass = 0;
	
	return 0;
}
