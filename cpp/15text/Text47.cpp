#include <iostream>
#include <fstream>
using namespace std;

int CharToInt(char C);

int main()
{
	//Task("Text47");
	const int stringSize = 50;

	char fileName[100];
	cin.getline(fileName, 100);
	ifstream file(fileName);

	char str[stringSize+1];
	int number, numbers = 0, sum = 0;
	bool isNegative, isDouble;
	while (file.peek() != -1)
	{
		file.getline(str, stringSize+1);
		
		int len = strlen(str);
		int i = 0;
		while ( (i < len) && (str[i] == ' ') )
			i++;

		isNegative = false;
		if (str[i] == '-')
		{
			isNegative = true;
			i++;
		}

		number = 0;
		isDouble = false;
		while ( (i < len) && (str[i] != ' ') )
		{
			if (str[i] == '.')
			{
				isDouble = true;
				break;
			}
			number = number * 10 + CharToInt(str[i]);
			i++;
		}

		if (isDouble)
			continue;
		
		if (isNegative)
			number *= -1;

		numbers++;
		sum += number;

	}
	file.close();

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