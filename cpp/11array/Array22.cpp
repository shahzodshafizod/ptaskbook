#include <iostream>
using namespace std;

int main()
{
	//Task("Array22");
	int N;
	cin >> N;
	double* mass = new double [N];
	double sum = 0;
	for (int i = 0; i < N; i++)
	{
		cin >> mass[i];
		sum += mass[i];
	}
	int K, L;
	cin >> K >> L;
	for (int i = K-1; i < L; i++)
		sum -= mass[i];
	cout << sum;
	delete [] mass;
	mass = 0;
	
	return 0;
}
