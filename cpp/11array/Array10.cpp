#include <iostream>
using namespace std;

int main()
{
	//Task("Array10");
	int N;
	cin >> N;
	
	int* mass = new int [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];
	
	for (int i = 0; i < N; i++)
		if (mass[i] % 2 == 0)
			cout << mass[i];
	
	for (int i = N-1; i >= 0; --i)
		if (mass[i] % 2 != 0)
			cout << mass[i];
	
	delete [] mass;
	mass = 0;
	
	return 0;
}
