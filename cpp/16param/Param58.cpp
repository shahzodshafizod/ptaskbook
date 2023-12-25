#include <iostream>
#include <fstream>
using namespace std;

void DecodeText(const char* S, int K);
void MoveToLeft(char& C, int K);

int main()
{
	//Task("Param58");
	int K;
	cin >> K;
	cin.ignore();
	char fileName[100];
	cin.getline(fileName, 100);
	DecodeText(fileName, K);
	
	return 0;
}

void DecodeText(const char* S, int K)
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
				MoveToLeft(str[i], K);

			tempFile << str << endl;
		}
		file.close();
		tempFile.close();

		remove(S);
		rename(tempFileName, S);
	}
}

void MoveToLeft(char& C, int K)
{
	int code = int(C);
	if ( (C >= 'а') && (C <= 'я') )
	{
		if (code-K < int('а'))
			code = int('я') - (K - (code - int('а') + 1));
		else
			code -= K;
	}
	else if ( (C >= 'А') && (C <= 'Я') )
	{
		if (code-K < int('А'))
			code = int('Я') - (K - (code - int('А') + 1));
		else
			code -= K;
	}
	C = char(code);
}