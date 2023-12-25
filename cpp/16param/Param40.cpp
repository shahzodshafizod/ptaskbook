#include <iostream>
using namespace std;

char* WordK(const char* S, int K);

int main()
{
	//Task("Param40");
	const int stringSize = 100;
	char S[stringSize+1];
	cin.getline(S, stringSize+1);
	int K;
	char* str = NULL;
	for (int i = 1; i < 4; i++)
	{
		cin >> K;
		str = WordK(S, K);
		cout << str;

		delete [] str;
		str = NULL;
	}
	
	return 0;
}

char* WordK(const char* S, int K)
{
	int len = strlen(S);
	int words = 0, from = 0, to = 0;
	for (int i = 0; i < len; i++)
	{
		while ( (i < len) && (S[i] == ' ') )
			i++;
		words++;
		from = i;
		while ( (i < len) && (S[i] != ' ') )
			i++;
		if (words == K)
		{
			to = i-1;
			break;
		}
	}
	int newLen = 0;
	if (to != 0)
		newLen = to - from + 1;
	char* theWord = new char [newLen+1];
	
	if (to != 0)
	{
		for (int i = 0, j = from; j <= to; i++, j++)
			theWord[i] = S[j];
	}
	theWord[newLen] = '\0';
	
	return theWord;
}