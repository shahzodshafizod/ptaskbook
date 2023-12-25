#include <iostream>
#include <fstream>
using namespace std;

bool isKitobat(char C);
char getBrother(char C);
void makeStr(const char* str, char** newStr, int index1, int index2);

int main()
{
	//Task("Text33");
	char C, c;
	cin >> C;
	cin.ignore();
	c = getBrother(C);
	const int stringSize = 80;
	char fileName[100], stringFileName[100], str[stringSize];
	cin.getline(fileName, 100);
	cin.getline(stringFileName, 100);
	ifstream file(fileName);
	ofstream stringFile(stringFileName, ios::binary);
	int index1, index2, strLen;
	char* newStr = NULL;
	bool hasC;
	while (true)
	{
		file.getline(str, stringSize);
		if (file.eof()) { file.clear(); break; }
		strLen = strlen(str);
		index1 = index2 = 0;
		for (int i = 0; i < strLen;)
		{
			while ( (i < strLen) && ( (str[i] == ' ') || isKitobat(str[i]) ) )
				i++;
			if (i == strLen)
				break;

			index1 = i;
			hasC = false;
			while ( (i < strLen) && (str[i] != ' ') && !isKitobat(str[i]) )
			{
				if ( (str[i] == C) || (str[i] == c) )
					hasC = true;
				i++;
			}
			index2 = i - 1;

			if (hasC)
			{
				makeStr(str, &newStr, index1, index2);
				stringFile.write(reinterpret_cast<const char*>(newStr), stringSize*sizeof(char));
			}
			if (i == strLen)
				break;

			i += (int)isKitobat(str[i]);
		}
	}
	file.close();
	stringFile.close();

	delete [] newStr;
	newStr = NULL;
	
	return 0;
}

void makeStr(const char* str, char** newStr, int index1, int index2)
{
	int len = index2 - index1 + 1;
	delete [] (*newStr);
	(*newStr) = new char [len+1];
	for (int i = 0, j = index1; j <= index2; i++, j++)
		(*newStr)[i] = str[j];
	(*newStr)[len] = '\0';
}

bool isKitobat(char C)
{
	bool result = (C == '.') || (C == ',') || (C == ':') || (C == ';') || (C == '\"') ||
				  (C == '!') || (C == '?') || (C == '-') || (C == '(') || (C == ')');
	return result;
}

char getBrother(char C)
{
	if ( (C >= 'А') && (C <= 'Я') )
	{
		int code = int(C) - int('А') + int('а');
		C = char(code);
	}
	else if ( (C >= 'а') && (C <= 'я') )
	{
		int code = int(C) - int('а') + int('А');
		C = char(code);
	}
	else if (C == 'Ё')
		C = 'ё';
	else if (C == 'ё')
		C = 'Ё';

	return C;
}