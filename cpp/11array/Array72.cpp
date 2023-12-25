#include <iostream>
using namespace std;

int main()
{
	//Task("Array72");
	
	int N;
	cin >> N;

	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	int K, L;
	cin >> K >> L;
	double temp;
	int from = K-1, to = L-1, howMany = to - from + 1;
	for (int i = from, j = 0; j < howMany/2; i++, j++)
	{
		temp = mass[i];
		mass[i] = mass[to-j];
		mass[to-j] = temp;
	}

	for (int i = 0; i < N; i++)
		cout << mass[i];

	delete [] mass;
	mass = 0;
	
	return 0;
}
