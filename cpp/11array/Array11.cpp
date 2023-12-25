#include <iostream>
using namespace std;

int main()
{
	//Task("Array11");
	int N;
	cin >> N;
	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];
	int K;
	cin >> K;
	int i = K - 1;
	while (i < N)
	{
		cout << mass[i];
		i += K;
	}
	delete [] mass;
	mass = 0;
	
	return 0;
}
