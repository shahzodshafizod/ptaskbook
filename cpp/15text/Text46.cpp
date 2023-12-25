#include <iostream>
#include <fstream>
using namespace std;

int CharToInt(char C);

int main()
{
	//Task("Text46");
	const int stringSize = 50;

	char fileInName[100], fileOutName[100];
	cin.getline(fileInName, 100);
	cin.getline(fileOutName, 100);

	ifstream fin(fileInName);
	ofstream fout(fileOutName, ios::binary);

	char str[stringSize+1];
	int len;
	double number;
	while (fin.peek() != -1)
	{
		fin.getline(str, stringSize+1);

		len = strlen(str);
		for (int i = 0; i < len; )
		{
			while ( (i < len) && (str[i] == ' ') )
				i++;
			bool isNegative = false;
			if (str[i] == '-')
			{
				isNegative = true;
				i++;
			}
			number = 0;
			while ( (i < len) && (str[i] != ' ') && (str[i] != '.') )
			{
				number = number * 10 + CharToInt(str[i]);
				i++;
			}
			if (str[i] == '.')
			{
				i++;
				int factor = 10;
				while ( (i < len) && (str[i] != ' ') )
				{
					number += CharToInt(str[i]) / (double)factor;
					factor *= 10;
					i++;
				}
				if (isNegative)
					number *= -1;
				fout.write(reinterpret_cast<const char*>(&number), sizeof(number));
			}
			else
			{
				while ( (i < len) && (str[i] != ' ') )
					i++;
			}
		}
	}
	
	return 0;
}

int CharToInt(char C)
{
	int farq = int(C) - int('0');
	
	if ( (farq >= 0) && (farq <= 9) )
		return farq;

	return -1;
}