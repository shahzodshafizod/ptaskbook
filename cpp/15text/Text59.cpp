#include <iostream>
#include <fstream>
using namespace std;

int CharToInt(char C);
void MoveRight(char& C, int K);
bool isRus(char C);

int main()
{
	//Task("Text59");
	const int stringSize = 100;

	char S[11], fileName[100], tempFileName[] = "tempFileName.txt";
	cin.getline(S, 11);
	cin.getline(fileName, 100);
	ifstream file(fileName);
	ofstream tempFile(tempFileName);

	int K[10];
	for (int i = 0; i < 10; i++)
		K[i] = CharToInt(S[i]);

	char str[stringSize+1];
	int len;
	while (file.peek() != -1)
	{
		file.getline(str, stringSize+1);
		len = strlen(str);
		for (int i = 0; i < len; i++)
			if (isRus(str[i]))
				MoveRight(str[i], K[i%10]);
		tempFile << str << endl;
	}
	file.close();
	tempFile.close();
	remove(fileName);
	rename(tempFileName, fileName);
	
	return 0;
}

int CharToInt(char C)
{
	int farq = int(C) - int('0');

	if ( (farq >= 0) && (farq <= 9) )
		return farq;

	return -1;
}

bool isRus(char C)
{
	return ( (C >= 'а') && (C <= 'я') || (C >= 'А') && (C <= 'Я') );
}

void MoveRight(char& C, int K)
{
	int code = int(C);

	if ( (C >= 'а') && (C <= 'я') )
	{
		if (code + K > int('я'))
			code = K - (int('я') - code + 1) + int('а');
		else
			code += K;
	}
	else if ( (C >= 'А') && (C <= 'Я') )
	{
		if (code + K > int('Я'))
			code = K - (int('Я') - code + 1) + int('А');
		else
			code += K;
	}

	C = char(code);
}