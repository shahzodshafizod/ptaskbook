#include <iostream>
using namespace std;

char GetNextChar(char C);

int main()
{
	//Task("String62");
	char str[1000];
	cin.getline(str, 1000);

	int index = 0;
	while (str[index] != '\0')
	{
		str[index] = GetNextChar(str[index]);
		index++;
	}

	cout << str;
	
	return 0;
}

char GetNextChar(char C)
{
	if (C == 'я')
		return 'а';
	if (C == 'Я')
		return 'А';
	if ( (C >= 'А') && (C < 'Я') || (C >= 'а') && (C < 'я') )
		return char(int(C) + 1);
	
	return C;
}