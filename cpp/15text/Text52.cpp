#include <iostream>
#include <fstream>
using namespace std;

int CharToInt(char C);

int main()
{
	//Task("Text52");
	const int stringSize = 100;

	char textFileName[100], intFileName[100];
	cin.getline(textFileName, 100);
	cin.getline(intFileName, 100);

	ifstream textFile(textFileName);
	ofstream intFile(intFileName, ios::binary);

	char str[stringSize+1];
	int len, number, factor, sum;
	char C;
	while (textFile.peek() != -1)
	{
		textFile.getline(str, stringSize+1);
		len = strlen(str);
		C = str[0];
		sum = 0;
		for (int i = 1; i < len; i++)
		{
			while ( (i < len) && (str[i] == ' ') )
				i++;
			
			factor = 1;
			if (str[i] == '-')
			{
				factor = -1;
				i++;
			}
			number = 0;
			while ( (i < len) && (str[i] != ' ') && (str[i] != C) )
			{
				number = number * 10 + CharToInt(str[i]);
				i++;
			}
			number *= factor;
			sum += number;

			while ( (i < len) && (str[i] != C) )
				i++;
		}
		intFile.write(reinterpret_cast<const char*>(&sum), sizeof(sum));
	}
	textFile.close();
	intFile.close();

	return 0;
}

int CharToInt(char C)
{
	int farq = int(C) - int('0');

	if ( (farq >= 0) && (farq <= 9) )
		return farq;

	return -1;
}