#include <iostream>
using namespace std;

int main()
{
	//Task("Minmax4");
	int N;
	cin >> N;

	double number, minimal;
	cin >> number;
	minimal = number;
	int index = 1;
	for (int i = 2; i <= N; i++)
	{
		cin >> number;

		if (number < minimal)
		{
			minimal = number;
			index = i;
		}
	}

	cout << index;
	
	return 0;
}
