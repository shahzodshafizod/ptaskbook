#include <iostream>
#include <fstream>
using namespace std;

int CharToInt(char C);

int main()
{
	//Task("Text50");
	const int stringSize = 80;

	char textFileName[100], stringFileName[100], doubleFileName[100];
	cin.getline(textFileName, 100);
	cin.getline(stringFileName, 100);
	cin.getline(doubleFileName, 100);

	ifstream textFile(textFileName);
	ofstream stringFile(stringFileName, ios::binary);
	ofstream doubleFile(doubleFileName, ios::binary);

	char str[stringSize + 1];
	int len;
	double number;
	while (textFile.peek() != -1)
	{
		textFile.getline(str, stringSize+1);
		len = strlen(str);
		int i = 30;
		bool isNegative = false;
		if (str[i] == '-')
		{
			isNegative = true;
			i++;
		}
		number = 0;
		for (; i < len; i++)
		{
			if (str[i] == '.')
				break;
			number = number * 10 + CharToInt(str[i]);
		}
		if (str[i] == '.')
		{
			i++;
			int factor = 10;
			while ( (i < len) )
			{
				number += CharToInt(str[i]) / (double)factor;
				factor *= 10;
				i++;
			}
		}
		if (isNegative)
			number *= -1;
		doubleFile.write(reinterpret_cast<const char*>(&number), sizeof(number));
		str[30] = '\0';
		stringFile.write(reinterpret_cast<const char*>(&str), sizeof(char)*stringSize);
	}
	textFile.close();
	stringFile.close();
	doubleFile.close();

	return 0;
}

int CharToInt(char C)
{
	int farq = int(C) - int('0');

	if ( (farq >= 0) && (farq <= 9) )
		return farq;

	return -1;
}