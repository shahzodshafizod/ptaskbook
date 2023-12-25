#include <iostream>
using namespace std;

int main()
{
	//Task("Minmax16");
	int N;
	cin >> N;

	int number, minimal;
	cin >> number;
	minimal = number;
	int counts = 0;

	for (int i = 2; i <= N; i++)
	{
		cin >> number;

		if (number < minimal)
		{
			minimal = number;
			counts = i-1;
		}
	}

	cout << counts;
	
	return 0;
}
