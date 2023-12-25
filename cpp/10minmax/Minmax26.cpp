#include <iostream>
using namespace std;

int main()
{
	//Task("Minmax26");
	int N;
	cin >> N;

	int number, counts = 0, prevIndex = 0, maxCounts = 0;
	for (int i = 1; i <= N; i++)
	{
		cin >> number;
		if (number % 2 == 0)
		{
			if (prevIndex == 0)
				counts = 1;
			
			else if (i-1 == prevIndex)
				counts++;
			
			else
			{
				if (maxCounts == 0)
					maxCounts = counts;
				else if (counts > maxCounts)
					maxCounts;
				counts = 1;
			}

			prevIndex = i;
		}
	}
	if (maxCounts == 0)
		maxCounts = counts;
	else if (counts > maxCounts)
		maxCounts = counts;

	cout << maxCounts;
	
	return 0;
}
