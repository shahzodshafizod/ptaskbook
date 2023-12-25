#include <iostream>
using namespace std;

int main()
{
	//Task("Array104");
	
	int N;
	cin >> N;

	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	int K, M;
	cin >> K >> M;
	
	double* temp = new double [N+M];
	int index = 0;
	for (int i = 0; i < N; i++)
	{
		if (i == K-1)
			for (int j = 0; j < M; j++)
				temp[index++] = 0;
		
		temp[index++] = mass[i];
	}

	delete [] mass;
	mass = temp;
	temp = 0;
	N += M;

	for (int i = 0; i < N; i++)
		cout << mass[i];

	delete [] mass;
	mass = 0;
	
	return 0;
}
