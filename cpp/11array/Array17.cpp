#include <iostream>
using namespace std;

int main()
{
	//Task("Array17");
	int N;
	cin >> N;
	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	int count = 0;
	int i = 0;
	while (count < N)
	{
		cout << mass[i]; count++;

		if (count >= N) break;
		
		cout << mass[i+1]; count++;

		if (count >= N) break;

		cout << mass[N-1-i]; count++;

		if (count >= N) break;

		cout << mass[N-2-i]; count++;

		i += 2;
	}

	delete [] mass;
	mass = 0;
	
	return 0;
}
