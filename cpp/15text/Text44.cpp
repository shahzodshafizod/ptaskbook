#include <iostream>
#include <fstream>
using namespace std;

int CharToInt(char C);

int main()
{
	//Task("Text44");
	const int stringSize = 50;

	char fileName[100];
	cin.getline(fileName, 100);

	ifstream file(fileName);
	int number, sum = 0;

	char str[stringSize+1];
	int numbers = 0;
	while (file.peek() != -1)
	{
		file.getline(str, stringSize+1);
		int len = strlen(str);

		int i = 0;
		while ( (i < len) && (str[i] == ' ') )
			i++;
		bool isNegative = false;
		if (str[i] == '-')
		{
			isNegative = true;
			i++;
		}
		number = 0;
		while ( (i < len) && (str[i] != ' ') )
		{
			number = number * 10 + CharToInt(str[i]);
			i++;
		}
		if (isNegative)
			number *= -1;

		numbers++;
		sum += number;
	}
	
	cout << numbers << sum;
	
	return 0;
}

int CharToInt(char C)
{
	int farq = int(C) - int('0');
	
	if ( (farq >= 0) && (farq <= 9) )
		return farq;
	
	return -1;
}