#include <iostream>
using namespace std;

int main()
{
	//Task("Array5");
	int N;
	cin >> N;
	int* mass = new int [N];
	mass[0] = mass[1] = 1;
	for (int i = 2; i < N; i++)
		mass[i] = mass[i-1] + mass[i-2];
	for (int i = 0; i < N; ++i)
		cout << mass[i];
	delete [] mass;
	mass = 0;
	
	return 0;
}
