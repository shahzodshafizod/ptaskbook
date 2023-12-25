#include <iostream>
#include <fstream>
using namespace std;

int CharToInt(char C);

int main()
{
	//Task("Text51");
	const int stringSize = 80;

	char textFileName[100], OutName1[100], OutName2[100], OutName3[100];
	cin.getline(textFileName, 100);
	cin.getline(OutName1, 100);
	cin.getline(OutName2, 100);
	cin.getline(OutName3, 100);

	ifstream textFile(textFileName);
	ofstream Out1(OutName1, ios::binary);
	ofstream Out2(OutName2, ios::binary);
	ofstream Out3(OutName3, ios::binary);

	double number;
	char str[stringSize+1];
	int len, fileNo;
	bool isNegative;
	while (textFile.peek() != -1)
	{
		textFile.getline(str, stringSize+1);
		
		len = strlen(str);
		fileNo = 0;
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
			while ( (i < len) && (str[i] != '.') )
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
			}
			if (isNegative)
				number *= -1;

			fileNo++;
			switch (fileNo)
			{
				case 1: Out1.write(reinterpret_cast<const char*>(&number), sizeof(number)); break;
				case 2: Out2.write(reinterpret_cast<const char*>(&number), sizeof(number)); break;
				case 3: Out3.write(reinterpret_cast<const char*>(&number), sizeof(number)); break;
			}
		}
	}
	textFile.close();
	Out1.close();
	Out2.close();
	Out3.close();

	return 0;
}

int CharToInt(char C)
{
	int farq = int(C) - int('0');

	if ( (farq >= 0) && (farq <= 9) )
		return farq;

	return -1;
}