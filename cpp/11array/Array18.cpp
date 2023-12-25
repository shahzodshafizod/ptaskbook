#include <iostream>
using namespace std;

int main()
{
	//Task("Array18");
	int mass[10];
	for (int i = 0; i < 10; i++)
		cin >> mass[i];

	int index = -1;
	for (int i = 0; i < 9; i++)
		if (mass[i] < mass[9])
		{
			index = i;
			break;
		}
	if (index < 0)
		cout << 0;
	else
		cout << mass[index];
	
	return 0;
}
