#include <iostream>
using namespace std;

int main()
{
	//Task("Minmax15");
	double B, C;
	cin >> B >> C;

	double number, maximal = 0;
	int index = -1;
	
	for (int i = 0; i < 10; i++)
	{
		cin >> number;

		if ( (number > B) && (number < C) )
		{
			if (maximal == 0)
			{
				maximal = number;
				index = i;
			}
			else if (number > maximal)
			{
				maximal = number;
				index = i;
			}
		}
	}

	cout << maximal << (index+1);
	
	return 0;
}
