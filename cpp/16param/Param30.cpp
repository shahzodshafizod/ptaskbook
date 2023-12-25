#include <iostream>
using namespace std;

bool isDigit(char C);
bool isLatin(char C);
int IsIdent(const char* S);

int main()
{
	//Task("Param30");
	const int stringSize = 100;
	char str[stringSize+1];
	for (int i = 1; i < 6; i++)
	{
		cin.getline(str, stringSize+1);
		cout << IsIdent(str);
	}
	
	return 0;
}

int IsIdent(const char* S)
{
	int len = strlen(S);
	if (len == 0)
		return -1;
	if (isDigit(S[0]))
		return -2;
	int errorIndex = 0;
	for (int i = 0; i < len; i++)
	{
		if ( (errorIndex == 0) && (S[i] != '_') && !isDigit(S[i]) && !isLatin(S[i]) )
		{
			errorIndex = i+1;
			break;
		}
	}
	return errorIndex;
}

bool isDigit(char C)
{
	return ( (C >= '0') && (C <= '9') );
}

bool isLatin(char C)
{
	return ( (C >= 'A') && (C <= 'Z') || (C >= 'a') && (C <= 'z') );
}