#include <iostream>
#include <fstream>
using namespace std;

char GetBrother(char);
void makeStr(const char* str, char** newStr, int index1, int index2);
bool isKitobat(char C);

int main()
{
	//Task("Text32");
	const int stringSize = 80;
	char C, c;
	cin >> C;
	cin.ignore();
	c = GetBrother(C);
	char fileName[100], stringFileName[100], str[stringSize];
	cin.getline(fileName, 100);
	cin.getline(stringFileName, 100);
	ifstream file(fileName);
	ofstream stringFile(stringFileName);
	int index1, index2, strLen;
	char* newStr = NULL;
	while (file.peek() != -1)
	{
		file.getline(str, stringSize);
		strLen = strlen(str);
		index1 = index2 = 0;
		for (int i = 0; i < strLen;)
		{
			while ( (i < strLen) && ( (str[i] == ' ') || isKitobat(str[i]) ) )
				i++;
			if (i == strLen)
				break;
			
			index1 = i;
			while ( (i < strLen) && (str[i] != ' ') && !isKitobat(str[i]) )
				i++;
			index2 = i-1;
			
			if ( (str[index1] == C) || (str[index1] == c) )
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
	for (int i = index1, j = 0; i <= index2; i++, j++)
		(*newStr)[j] = str[i];
	(*newStr)[len] = '\0';
}

char GetBrother(char C)
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

bool isKitobat(char C)
{
	bool result = (C == '!') || (C == '?') || (C == '.') || (C == ',') || (C == '\"') || 
				  (C == '-') || (C == '(') || (C == ')') || (C == ':') || (C == ';');
	return result;
}