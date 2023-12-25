#include <iostream>
using namespace std;

int main()
{
	//Task("Array50");
	int N;
	cin >> N;
	int* mass = new int [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];

	int count = 0;
	for (int i = 0; i < N-1; i++)
		for (int j = i+1; j < N; j++)
			if (mass[i] > mass[j])
				count++;
	
	cout << count;

	delete [] mass;
	mass = 0;
	
	return 0;
}
