#include <iostream>
using namespace std;

int main()
{
	//Task("Series1");
	double number, sum = 0;
	for (int i = 1; i <= 10; i++)
	{
		cin >> number;
		sum += number;
	}
	cout << sum;
	
	return 0;
}
