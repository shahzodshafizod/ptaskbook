#include <iostream>
using namespace std;

int DigitCount(const char* S, int index = 0);
bool isDigit(char C);

int main()
{
	//Task("Recur12");
	const int stringSize = 100;
	char str[stringSize+1];
	for (int i = 0; i < 5; i++)
	{
		cin.getline(str, stringSize+1);
		cout << DigitCount(str);
	}
	return 0;
}

int DigitCount(const char* S, int index)
{
	if (S[index] == '\0')
		return 0;

	return (int)isDigit(S[index]) + DigitCount(S, index+1);
}

bool isDigit(char C)
{
	return ( (C >= '0') && (C <= '9') );
}