#include <iostream>
using namespace std;

char* DecToHex(int number);
char toHexDigit(int number);

int main()
{
	//Task("Param45");
	int number;
	char* str = NULL;
	for (int i = 1; i < 6; i++)
	{
		cin >> number;
		str = DecToHex(number);
		cout << str;

		delete [] str;
		str = NULL;
	}
	
	return 0;
}

char* DecToHex(int number)
{
	int digits = number == 0 ? 1 : 0, temp = number;
	while (temp > 0)
	{
		temp /= 16;
		digits++;
	}
	char* str = new char [digits+1];
	int index = 0;
	if (number == 0)
		str[index++] = '0';
	else
		while (number > 0)
		{
			str[index++] = toHexDigit(number % 16);
			number /= 16;
		}
	str[index--] = '\0';

	for (int i = 0; i < index; i++, index--)
	{
		char C = str[i];
		str[i] = str[index];
		str[index] = C;
	}

	return str;
}

char toHexDigit(int number)
{
	if ( (number >= 0) && (number <= 9) )
		return char(int('0') + number);
	
	number -= 10;
	return char( int('A') + number );
}