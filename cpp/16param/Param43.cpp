#include <iostream>
using namespace std;

char* DecompressStr(const char* S);

int main()
{
	//Task("Param43");
	const int stringSize = 100;
	char S[stringSize+1], *str = NULL;
	for (int i = 1; i < 6; i++)
	{
		cin.getline(S, 101);
		str = DecompressStr(S);
		cout << str;

		delete [] str;
		str = NULL;
	}
	
	return 0;
}

int CharToInt(char C)
{
	return int(C) - int('0');
}

char* DecompressStr(const char* S)
{
	int len = strlen(S);
	char* str = new char [100];
	int index = 0;
	for (int i = 0; i < len; i++)
	{
		if (S[i] == '{')
		{
			char C = S[i-1];
			i++;
			int Cs = 0;
			while (S[i] != '}')
				Cs = Cs*10 + CharToInt(S[i++]);
			for (int j = 1; j < Cs; j++)
				str[index++] = C;
		}
		else
			str[index++] = S[i];
	}
	str[index] = '\0';

	return str;
}