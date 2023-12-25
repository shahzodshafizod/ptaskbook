#include <iostream>
using namespace std;

int main()
{
	//Task("Array47");
	int N;
	cin >> N;
	int* mass = new int [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	int counts[100] = {0};
	for (int i = 0; i < N; i++)
		counts[mass[i]]++;

	int count = 0;
	for (int i = 0; i < 100; i++)
		if (counts[i] != 0)
			count++;
	
	cout << count;

	delete [] mass;
	mass = 0;
	
	return 0;
}
