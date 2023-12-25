#include <iostream>
using namespace std;

int main()
{
	//Task("Array23");
	int N;
	cin >> N;
	double* mass = new double [N];
	double sum = 0;
	for (int i = 0; i < N; i++)
	{
		cin >> mass[i];
		sum += mass[i];
	}
	int count = N;
	int K, L;
	cin >> K >> L;
	for (int i = K-1; i < L; i++)
	{
		sum -= mass[i];
		count--;
	}
	cout << sum / count;
	delete [] mass;
	mass = 0;
	
	return 0;
}
