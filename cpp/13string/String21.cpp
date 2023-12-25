#include <iostream>
using namespace std;

int main()
{
	//Task("String21");
	unsigned int number;
	cin >> number;
	while (number > 0)
	{
		cout << char( int('0') + number%10 );
		number /= 10;
	}
	
	return 0;
}
