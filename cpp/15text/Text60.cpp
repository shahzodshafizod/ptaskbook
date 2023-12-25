#include <iostream>
#include <fstream>
using namespace std;

struct Position
{
	int pos;
	bool hasValue;
};

bool isRus(char C);
void MoveLeft(char& C, int K);

int main()
{
	//Task("Text60");
	const int stringSize = 100;
	
	char S[stringSize+1];
	cin.getline(S, stringSize+1);
	int strLen = strlen(S);

	char fileName[100];
	cin.getline(fileName, 100);
	ifstream file(fileName);
	
	char str[stringSize+1];
	file.getline(str, stringSize+1);
	int len = strlen(str);

	Position K[10];
	for (int i = 0; i < 10; i++)
		K[i].hasValue = false;
	
	bool canMove = false;
	if (len == strLen)
	{
		canMove = true;
		int temp;
		for (int i = 0; i < len; i++)
			if ( isRus(S[i]) && isRus(str[i]) )
			{
				temp = int(str[i]) - int(S[i]);
				temp += temp < 0 ? 32 : 0;
				
				if (K[i%10].hasValue)
				{
					if (K[i%10].pos != temp)
					{
						canMove = false;
						break;
					}
				}
				else
				{
					K[i%10].pos = temp;
					K[i%10].hasValue = true;
				}
			}

		for (int i = 0; i < 10; i++)
			if (K[i].hasValue == false)
			{
				canMove = false;
				break;
			}
	}

	if (canMove)
	{
		char tempFileName[] = "tempFileName.txt";
		ofstream tempFile(tempFileName);
		tempFile << S;

		while (file.peek() != -1)
		{
			file.getline(str, stringSize+1);
			len = strlen(str);
			for (int i = 0; i < len; i++)
				if (isRus(str[i]))
					MoveLeft(str[i], K[i%10].pos);
			tempFile << endl << str;
		}
		file.close();
		tempFile.close();
		remove(fileName);
		rename(tempFileName, fileName);

	}
	else
		file.close();

	return 0;
}

bool isRus(char C)
{
	return ( (C >= 'а') && (C <= 'я') || (C >= 'А') && (C <= 'Я') );
}

void MoveLeft(char& C, int K)
{
	int code = int(C);

	if ( (C >= 'а') && (C <= 'я') )
	{
		if (code - K < int('а'))
			code = int('я') - (K - (code - int('а') + 1));
		else
			code -= K;
	}
	else if ( (C >= 'А') && (C <= 'Я') )
	{
		if (code - K < int('А'))
			code = int('Я') - (K - (code - int('А') + 1));
		else
			code -= K;
	}

	C = char(code);
}