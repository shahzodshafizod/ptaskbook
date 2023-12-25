#include <iostream>
using namespace std;

int main()
{
	//Task("Array43");
	
	int N;
	cin >> N;
	
	int* mass = new int [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	int count = 1;
	for (int i = 1; i < N; i++)
		if (mass[i] != mass[i-1])
			count++;
	
	cout << count;

	delete [] mass;
	mass = 0;
	
	return 0;
}
