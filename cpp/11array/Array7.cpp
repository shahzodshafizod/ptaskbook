#include <iostream>
using namespace std;

int main()
{
	//Task("Array7");
	int N;
	cin >> N;
	double* mass = new double [N];
	for (int i = 0; i < N; ++i)
		cin >> mass[i];
	for (int i = N-1; i >= 0; --i)
		cout << mass[i];
	delete [] mass;
	mass = 0;
	
	return 0;
}
