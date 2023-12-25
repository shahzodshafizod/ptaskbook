#include <iostream>
using namespace std;

int main()
{
	//Task("Minmax29");
	int N;
	cin >> N;

	int number, minimal, counts = 1, maxCounts = 0;
	cin >> number;
	minimal = number;
	for (int i = 2; i <= N; i++)
	{
		cin >> number;
		if (number == minimal)
			counts++;

		else if (number < minimal)
		{
			minimal = number;
			counts = 1;
			maxCounts = 0;
		}

		else
		{
			if (maxCounts == 0)
				maxCounts = counts;
			else if (counts > maxCounts)
				maxCounts = counts;
			counts = 0;
		}
	}
	if (maxCounts == 0)
		maxCounts = counts;
	else if (counts > maxCounts)
		maxCounts = counts;

	cout << maxCounts;
	
	return 0;
}
