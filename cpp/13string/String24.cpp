#include <iostream>
using namespace std;

int main()
{
	//Task("String24");
	char str[100];
	cin.getline(str, 100);

	int len = 0;
	while (str[len] != '\0')
		len++;

	int du = 1, result = 0;
	for (int i = len-1; i >= 0; i--)
	{
		if (str[i] == '1')
			result += du;
		du *= 2;
	}

	int temp = 0, digits = 0;
	while (result > 0)
	{
		temp = temp * 10 + result % 10;
		result /= 10;
		digits++;
	}

	char* newStr = new char [digits+1];
	for (int i = 0; i < digits; i++)
	{
		newStr[i] = char( int('0') + temp % 10);
		temp /= 10;
	}
	newStr[digits] = '\0';

	cout << newStr;

	delete [] newStr;
	newStr = NULL;
	
	return 0;
}
