#include <iostream>
#include <fstream>
using namespace std;

struct Char
{
	int code;
	int counts;
};

void sortKam(Char* chars, int aSize);
char IntToChar(int d);

int main()
{
	//Task("Text58");
	const int stringSize = 80, charCounts = 256;

	char textFileName[100], stringFileName[100];
	cin.getline(textFileName, 100);
	cin.getline(stringFileName, 100);

	ifstream textFile(textFileName);
	ofstream stringFile(stringFileName, ios::binary);

	int len, code;
	int chars[charCounts] = {0};
	char str[stringSize+1];
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

	int resultSize = to - from + 1;
	Char* ramzho = new Char [resultSize];
	
	for (int i = from, j = 0; i <= to; i++, j++)
	{
		ramzho[j].code = i;
		ramzho[j].counts = chars[i];
	}

	sortKam(ramzho, resultSize);

	char C;
	for (int i = 0; i < resultSize; i++)
		if (ramzho[i].counts != 0)
		{
			int j = 0;
			str[j++] = char(ramzho[i].code);
			str[j++] = '-';
			while (ramzho[i].counts > 0)
			{
				str[j++] = IntToChar(ramzho[i].counts % 10);
				ramzho[i].counts /= 10;
			}

			int darozi = j - 2;
			for (int k = 0, l = 2; k < darozi/2; k++, l++)
			{
				C = str[l];
				str[l] = str[2 + darozi - 1 - k];
				str[2 + darozi - 1 - k] = C;
			}
			str[j] = '\0';
			stringFile.write(reinterpret_cast<const char*>(&str), sizeof(char)*stringSize);
		}

	stringFile.close();

	return 0;
}

void sortKam(Char* chars, int aSize)
{
	Char temp;
	for (int i = 0; i < aSize; i++)
	{
		for (int j = 1; j < aSize-i; j++)
		{
			if (chars[j].counts > chars[j-1].counts)
			{
				temp = chars[j];
				chars[j] = chars[j-1];
				chars[j-1] = temp;
			}
			else if (chars[j].counts == chars[j-1].counts)
			{
				if (chars[j].code < chars[j-1].code)
				{
					temp = chars[j];
					chars[j] = chars[j-1];
					chars[j-1] = temp;
				}
			}
		}
	}
}

char IntToChar(int d)
{
	int code = int('0') + d;
	return char(code);
}