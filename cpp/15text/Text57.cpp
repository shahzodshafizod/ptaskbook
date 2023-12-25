#include <iostream>
#include <fstream>
using namespace std;

char IntToChar(int d);

int main()
{
	//Task("Text57");
	const int stringSize = 80, charCounts = 256;

	char textFileName[100], stringFileName[100];
	cin.getline(textFileName, 100);
	cin.getline(stringFileName, 100);

	ifstream textFile(textFileName);
	ofstream stringFile(stringFileName, ios::binary);

	char str[stringSize+1];

	int chars[charCounts] = {0};
	int len, code;
	while (textFile.peek() != -1)
	{
		textFile.getline(str, stringSize+1);
		len = strlen(str);
		for (int i = 0; i < len; i++)
		{
			code = int(str[i]);
			code += code < 0 ? 256 : 0;
			chars[code]++;
		}
	}
	textFile.close();

	int from = int('а');
	from += from < 0 ? 256 : 0;
	int to = int('я');
	to += to < 0 ? 256 : 0;

	for (int i = from; i <= to; i++)
	{
		if (chars[i] != 0)
		{
			int j = 0;
			str[j++] = char(i);
			str[j++] = '-';

			while (chars[i] > 0)
			{
				str[j++] = IntToChar(chars[i] % 10);
				chars[i] /= 10;
			}
			
			int darozi = j - 2;
			char C;
			for (int k = 0, l = 2; k < darozi/2; k++, l++)
			{
				C = str[l];
				str[l] = str[2+darozi-1-k];
				str[2+darozi-1-k] = C;
			}
			
			str[j] = '\0';
			stringFile.write(reinterpret_cast<const char*>(&str), sizeof(char)*stringSize);
		}
	}
	
	stringFile.close();

	return 0;
}

char IntToChar(int d)
{
	int code = int('0') + d;
	return char(code);
}