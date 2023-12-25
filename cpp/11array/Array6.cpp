#include <iostream>
using namespace std;

int main()
{
	//Task("Array6");
	int N;
	cin >> N;
	int A, B;
	cin >> A >> B;
	int* mass = new int [N];
	mass[0] = A;
	mass[1] = B;
	int sum;
	for (int i = 2; i < N; i++)
	{
		sum = 0;
		for (int j = 0; j < i; j++)
			sum += mass[j];
		mass[i] = sum;
	}
	for (int i = 0; i < N; ++i)
		cout << mass[i];
	delete [] mass;
	mass = 0;
	
	return 0;
}
