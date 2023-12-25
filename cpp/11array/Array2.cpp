#include <iostream>
using namespace std;

int main()
{
	//Task("Array2");
	int N;
	cin >> N;
	int* mass = new int [N];
	mass[0] = 2;
	for (int i = 1; i < N; ++i)
		mass[i] = mass[i-1] * 2;
	for (int i = 0; i < N; i++)
		cout << mass[i];
	delete [] mass;
	mass = 0;
	
	return 0;
}
