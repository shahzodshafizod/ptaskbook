#include <iostream>
using namespace std;

int main()
{
	//Task("Series13");
	int number, sum = 0;
	do
	{
		cin >> number;
		if (number == 0)
			break;
		if ( (number > 0) && (number % 2 == 0) )
			sum += number;
	}
	while (true);
	cout << sum;
	
	return 0;
}
