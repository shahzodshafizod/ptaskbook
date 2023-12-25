#include <iostream>
using namespace std;

int main()
{
	//Task("Array101");
	
	int N;
	cin >> N;

	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	int K;
	cin >> K;
	int index = 0;
	double* temp = new double [N+1];
	for (int i = 0; i < N; i++)
	{
		if (i == K-1)
			temp[index++] = 0;
		
		temp[index++] = mass[i];
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
