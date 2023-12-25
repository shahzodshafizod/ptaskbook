#include <iostream>
#include <fstream>
using namespace std;

void UpperLowerCase(char& C);

int main()
{
	//Task("Text19");
	const int stringSize = 80;
	char fileName[100], str[stringSize], tempFileName[] = "tempFileName.txt";
	cin.getline(fileName, 100);
	ifstream file(fileName);
	ofstream tempFile(tempFileName);
	
	int length;
	while (true)
	{
		file.getline(str, stringSize);
		if (file.eof())
		{
			file.clear();
			break;
		}
		length = strlen(str);
		for (int i = 0; i < length; i++)
			UpperLowerCase(str[i]);
		tempFile << str << endl;
	}
	file.close();
	tempFile.close();
	remove(fileName);
	rename(tempFileName, fileName);
	
	return 0;
}

void UpperLowerCase(char& C)
{
	if ( (C >= 'А') && (C <= 'Я') )
	{
		int code = int(C) - int('А') + int('а');
		C = char(code);
	}
	else if ( (C >= 'а') && (C <= 'я') )
	{
		int code = (C) - int('а') + int('А');
		C = char(code);
	}
	else if (C == 'Ё')
		C = 'ё';
	else if (C == 'ё')
		C = 'Ё';
}