#include <iostream>
using namespace std;

int main()
{
	//Task("String20");
	unsigned int number, temp;
	cin >> number;
	temp = 0;
	int digits = 0;
	while (number > 0)
	{
		temp = temp * 10 + number % 10;
		number /= 10;
		digits++;
	}
	for (int i = 1; i <= digits; i++)
	{
		cout << char( int('0') + temp%10 );
		temp /= 10;
	}
	
	return 0;
}