#include <iostream>
using namespace std;

int main()
{
	//Task("Array48");
	int N;
	cin >> N;
	int* mass = new int [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	int counts[100] = {0};
	for (int i = 0; i < N; i++)
		counts[mass[i]]++;
	
	int maximal = counts[0];
	for (int i = 1; i < 100; i++)
		if (counts[i] > maximal)
			maximal = counts[i];
	
	cout << maximal;
	
	delete [] mass;
	mass = 0;
	
	return 0;
}
