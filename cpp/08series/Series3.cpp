#include <iostream>
using namespace std;

int main()
{
	//Task("Series3");
	double number, sum = 0;
	for (int i = 1; i <= 10; i++)
	{
		cin >> number;
		sum += number;
	}
	cout << sum / 10;
	
	return 0;
}
