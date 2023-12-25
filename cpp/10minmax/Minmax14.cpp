#include <iostream>
using namespace std;

int main()
{
	//Task("Minmax14");
	double B;
	cin >> B;

	double number, minimal = 0;
	int index = -1;
	
	for (int i = 0; i < 10; i++)
	{
		cin >> number;

		if (number > B)
		{
			if (minimal == 0)
			{
				minimal = number;
				index = i;
			}
			else if (number < minimal)
			{
				minimal = number;
				index = i;
			}
		}
	}
	
	cout << minimal << (index+1);
	
	return 0;
}
