#include <iostream>
#include <fstream>
using namespace std;

int CharToInt(char C);

int main()
{
	//Task("Text48");
	const int stringSize = 50;

	char fileName[100], fileOutName[100];
	cin.getline(fileName, 100);
	cin.getline(fileOutName, 100);

	ifstream file(fileName);
	ofstream fout(fileOutName, ios::binary);

	char str[stringSize+1];

	int number, len;
	bool isDouble, isNegative;
	while (file.peek() != -1)
	{
		file.getline(str, stringSize+1);

		len = strlen(str);

		for (int i = 0; i < len; )
		{
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
					i++;
					break;
				}
				number = number * 10 + CharToInt(str[i]);
				i++;
			}
			if (isDouble)
			{
				while ( (i < len) && (str[i] != ' ') )
					i++;
				continue;
			}
			if (isNegative)
				number *= -1;
			fout.write(reinterpret_cast<const char*>(&number), sizeof(number));
		}
	}
	fout.close();
	file.close();

	return 0;
}

int CharToInt(char C)
{
	int farq = int(C) - int('0');
	if ( (farq >= 0) && (farq <= 9) )
		return farq;

	return -1;
}