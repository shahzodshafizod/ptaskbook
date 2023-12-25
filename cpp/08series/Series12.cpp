#include <iostream>
using namespace std;

int main()
{
	//Task("Series12");
	int number, counts = 0;
	do
	{
		cin >> number;
		if (number == 0)
			break;
		counts++;
	}
	while (true);
	cout << counts;
	
	return 0;
}
