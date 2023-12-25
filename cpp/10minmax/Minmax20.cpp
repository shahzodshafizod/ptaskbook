#include <iostream>
using namespace std;

int main()
{
	//Task("Minmax20");
	int N;
	cin >> N;

	int number, minimal, maximal, mins = 1, maxs = 1;
	cin >> number;
	minimal = maximal = number;
	for (int i = 2; i <= N; i++)
	{
		cin >> number;

		if (number == maximal)
			maxs++;

		if (number == minimal)
			mins++;

		if (number < minimal)
		{
			minimal = number;
			mins = 1;
		}

		if (number > maximal)
		{
			maximal = number;
			maxs = 1;
		}
	}

	int counts = (mins == N) ? mins : mins + maxs;
	cout << counts;
	
	return 0;
}
