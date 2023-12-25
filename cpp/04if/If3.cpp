#include <iostream>
using namespace std;

int main()
{
	//Task("If3");
	int number;
	cin >> number;
	if (number > 0)
		number -= 8;
	else if (number < 0)
		number += 6;
	else
		number = 10;
	cout << number;
	
	return 0;
}
