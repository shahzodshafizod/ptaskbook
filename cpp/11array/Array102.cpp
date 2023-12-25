#include <iostream>
using namespace std;

int main()
{
	//Task("Array102");
	
	int N;
	cin >> N;

	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	int K;
	cin >> K;
	double* temp = new double [N+1];
	int index = 0;
	for (int i = 0; i < N; i++)
	{
		temp[index++] = mass[i];
		
		if (i == K-1)
			temp[index++] = 0;
	}
	
	delete [] mass;
	mass = temp;
	temp = 0;
	N++;
	
	for (int i = 0; i < N; i++)
		cout << mass[i];

	delete [] mass;
	mass = 0;
	
	return 0;
}
