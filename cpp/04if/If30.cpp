#include <iostream>
using namespace std;

int main()
{
	//Task("If30");
	int number;
	cin >> number;
	string str = "";
	
	if (number % 2 != 0)
		str += "не";
	str += "четное ";

	if (number <= 9)
		str += "одно";
	else if (number <= 99)
		str += "дву";
	else
		str += "трех";
	
	str += "значное число";
	
	cout << str;
	
	return 0;
}
