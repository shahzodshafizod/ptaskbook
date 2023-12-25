#include <iostream>
using namespace std;

int HexToDec(const char* S);
int HexDigitToInt(char C);

int main()
{
	//Task("Param47");
	const int stringSize = 100;
	char str[stringSize+1];
	for (int i = 0; i < 5; i++)
	{
		cin.getline(str, stringSize+1);
		cout << HexToDec(str);
	}
	
	return 0;
}

int HexToDec(const char* S)
{
	int len = 0;
	while (S[len] != '\0')
		len++;

	int degree = 1, result = 0;
	for (int i = len-1; i >= 0; i--)
	{
		result += HexDigitToInt(S[i]) * degree;
		degree *= 16;
	}
	
	return result;
}

int HexDigitToInt(char C)
{
	int number = int(C) - int('0');
	if ( (number >= 0) && (number <= 9) )
		return number;

	number = int (C) - int('A') + 10;
	
	return number;
}