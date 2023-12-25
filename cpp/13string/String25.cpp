#include <iostream>
using namespace std;

int main()
{
	//Task("String25");
	char str[100];
	cin.getline(str, 100);

	int len = 0, number = 0;
	while (str[len] != '\0')
	{
		number = number * 10 + ( int(str[len]) - int('0') );
		len++;
	}
	
	int temp = number, digits = 0;
	while (temp > 0)
	{
		temp /= 2;
		digits++;
	}

	char* newStr = new char [digits+1];
	for (int i = digits-1; i >= 0; i--)
	{
		newStr[i] = char( int('0') + number % 2);
		number /= 2;
	}
	newStr[digits] = '\0';

	cout << newStr;

	delete [] newStr;
	newStr = NULL;
	
	return 0;
}
