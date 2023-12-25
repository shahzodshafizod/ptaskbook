#include <iostream>
using namespace std;

int main()
{
	//Task("Array90");
	
	int N, K;
	cin >> N;

	double* mass = new double [N];
	double* temp = new double [N-1];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	cin >> K;
	int index = 0;
	for (int i = 0; i < K-1; i++)
		temp[index++] = mass[i];
	for (int i = K; i < N; i++)
		temp[index++] = mass[i];

	delete [] mass;
	mass = temp;
	temp = 0;

	N--;
	for (int i = 0; i < N; i++)
		cout << mass[i];

	delete [] mass;
	mass = 0;
	
	return 0;
}
