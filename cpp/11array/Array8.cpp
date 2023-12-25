#include <iostream>
using namespace std;

int main()
{
	//Task("Array8");
	int N;
	cin >> N;
	int* mass = new int [N];
	for (int i = 0; i < N; ++i)
		cin >> mass[i];
	int count = 0;
	for (int i = 0; i < N; i++)
		if (mass[i] % 2 != 0)
		{
			cout << mass[i];
			count++;
		}
	cout << count;
	delete [] mass;
	mass = 0;
	
	return 0;
}
