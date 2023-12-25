#include <iostream>
using namespace std;

int main()
{
	//Task("Array38");
	int N;
	cin >> N;
	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	int count = 0;
	bool regress = false;
	for (int i = 1; i < N; i++)	
	{
		if (mass[i] < mass[i-1])
			regress = true;
		
		else if (regress)
		{
			count++;
			regress = false;
		}
	}
	if (regress)
		count++;

	cout << count;

	delete [] mass;
	mass = 0;
	
	return 0;
}
