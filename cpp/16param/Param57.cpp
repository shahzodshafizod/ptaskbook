#include <iostream>
#include <fstream>
using namespace std;

void EncodeText(const char* S, int K);
void MoveToRight(char& C, int K);

int main()
{
	//Task("Param57");
	int K;
	cin >> K;
	cin.ignore();
	char fileName[100];
	cin.getline(fileName, 100);
	EncodeText(fileName, K);
	
	return 0;
}

void EncodeText(const char* S, int K)
{
	ifstream file(S);
	if (file.is_open())
	{
		char tempFileName[] = "tempFileName.txt";
		ofstream tempFile(tempFileName);
		const int stringSize = 100;
		char str[stringSize+1];
		while (file.peek() != -1)
		{
			file.getline(str, stringSize+1);

			int len = strlen(str);
			for (int i = 0; i < len; i++)
				MoveToRight(str[i], K);

			tempFile << str << endl;
		}
		file.close();
		tempFile.close();
		
		remove(S);
		rename(tempFileName, S);
	}
}

void MoveToRight(char& C, int K)
{
	int code = int(C);
	if ( (C >= 'а') && (C <= 'я') )
	{
		if (code+K > int('я'))
			code = (K - (int('я') - code + 1)) + int('а');
		else
			code += K;
	}
	else if ( (C >= 'А') && (C <= 'Я') )
	{
		if (code+K > int('Я'))
			code = (K - (int('Я') - code + 1)) + int('А');
		else
			code += K;
	}
	C = char(code);
}