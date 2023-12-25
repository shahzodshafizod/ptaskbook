#include <iostream>
using namespace std;

int main()
{
	//Task("Minmax8");
	int N;
	cin >> N;

	int number, minimal;
	cin >> number;
	minimal = number;
	int firstMinIndex, lastMinIndex;
	firstMinIndex = lastMinIndex = 1;

	for (int i = 2; i <= N; i++)
	{
		cin >> number;

		if (number == minimal)
			lastMinIndex = i;

		if (number < minimal)
		{
			minimal = number;
			firstMinIndex = lastMinIndex = i;
		}
	}

	cout << firstMinIndex << lastMinIndex;
	
	return 0;
}
