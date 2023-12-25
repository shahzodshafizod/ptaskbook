#include <iostream>
using namespace std;

char* DecToBin(int N);

int main()
{
	//Task("Param44");
	int number;
	char* str = NULL;
	for (int i = 1; i < 6; i++)
	{
		cin >> number;
		str = DecToBin(number);
		cout << str;
	}
	
	return 0;
}

char* DecToBin(int N)
{
	int digits = N == 0 ? 1 : 0, temp = N;
	while (temp > 0)
	{
		temp /= 2;
		digits++;
	}

	char* str = new char [digits+1];
	int index = 0;
	if (N == 0)
		str[index++] = '0';
	else
		while (N > 0)
		{
			str[index++] = char(int('0') + N % 2);
			N /= 2;
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