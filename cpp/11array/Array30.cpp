#include <iostream>
using namespace std;

int main()
{
	//Task("Array30");
	int N;
	cin >> N;
	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];
	int count = 0;
	for (int i = 0; i < N-1; i++)
		if (mass[i] > mass[i+1])
		{
			cout << i+1;
			count++;
		}
	cout << count;
	delete [] mass;
	mass = 0;
	
	return 0;
}
