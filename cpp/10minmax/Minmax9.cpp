#include <iostream>
using namespace std;

int main()
{
	//Task("Minmax9");
	int N;
	cin >> N;

	int number, maximal;
	int firstMaxIndex, lastMaxIndex;
	cin >> number;
	maximal = number;
	firstMaxIndex = lastMaxIndex = 1;

	for (int i = 2; i <= N; i++)
	{
		cin >> number;

		if (number == maximal)
			lastMaxIndex = i;

		if (number > maximal)
		{
			maximal = number;
			firstMaxIndex = lastMaxIndex = i;
		}
	}

	cout << firstMaxIndex << lastMaxIndex;
	
	return 0;
}
