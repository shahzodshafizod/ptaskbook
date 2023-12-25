#include <iostream>
using namespace std;

int main()
{
	//Task("Minmax13");
	int N;
	cin >> N;

	int number, maxOdd = 0, index = -1;
	for (int i = 0; i < N; i++)
	{
		cin >> number;

		if (number % 2 != 0)
		{
			if (maxOdd == 0)
			{
				maxOdd = number;
				index = i;
			}
			else if (number > maxOdd)
			{
				maxOdd = number;
				index = i;
			}
		}
	}

	cout << (index+1);
	
	return 0;
}
