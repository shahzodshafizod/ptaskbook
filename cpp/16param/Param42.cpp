#include <iostream>
using namespace std;

char* CompressStr(const char* S);

int main()
{
	//Task("Param42");
	const int stringSize = 100;
	char S[stringSize+1], *str = NULL;
	
	for (int i = 1; i < 6; i++)
	{
		cin.getline(S, stringSize+1);
		str = CompressStr(S);
		cout << str;
		delete [] str;
		str = NULL;
	}
	
	return 0;
}

char IntToChar(int digit)
{
	return char(int('0') + digit);
}

char* CompressStr(const char* S)
{
	int len = strlen(S);
	char C, *str = new char [len+1];
	int Cs, index = 0;
	for (int i = 0; i < len;)
	{
		C = S[i++];
		Cs = 1;
		while ( (i < len) && (C == S[i]) )
		{
			Cs++;
			i++;
		}
		if (Cs > 4)
		{
			str[index++] = C;
			str[index++] = '{';
			int to, from = index;
			while (Cs > 0)
			{
				str[index++] = IntToChar(Cs % 10);
				Cs /= 10;
			}
			to = index - 1;
			char temp;
			while (from < to)
			{
				temp = str[from];
				str[from] = str[to];
				str[to] = temp;
				
				from++;
				to--;
			}
			str[index++] = '}';
		}
		else
			for (int j = 0; j < Cs; j++)
				str[index++] = C;
	}
	str[index] = '\0';

	return str;
}