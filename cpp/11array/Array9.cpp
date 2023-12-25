#include <iostream>
using namespace std;

int main()
{
	//Task("Array9");
	int N;
	cin >> N;
	int* mass = new int [N];
	for (int i = 0; i < N; ++i)
		cin >> mass[i];
	int count = 0;
	for (int i = N-1; i >= 0; --i)
		if (mass[i] % 2 == 0)
		{
			cout << mass[i];
			count++;
		}
	cout << count;
	
	return 0;
}
