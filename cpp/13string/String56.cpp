#include <iostream>
using namespace std;

struct Word
{
	int index1;
	int index2;
	int length;
};

bool isSymbol(const char* symb, int len, char C);

int main()
{
	//Task("String56");
	char symb[] = "!?().,;:-\"";
	int arraySize = sizeof(symb) / sizeof(symb[0]);
	
	char str[1000];
	cin.getline(str, 1000);

	int len = 0;
	while (str[len] != '\0')
		len++;

	Word word, minWord;
	minWord.length = -1;
	for (int i = 0; i < len; )
	{
		while ( (i < len) && ( (str[i] == ' ') || isSymbol(symb, arraySize, str[i]) ) )
			i++;
		if (i == len)
			break;

		word.index1 = i;
		word.length = 0;
		while ( (i < len) && (str[i] != ' ') && !isSymbol(symb, arraySize, str[i]) )
		{
			word.length++;
			i++;
		}
		word.index2 = i-1;

		if (minWord.length < 0)
			minWord = word;
		else if (word.length <= minWord.length)
			minWord = word;
	}

	char* newStr = new char [minWord.length + 1];
	for (int i = 0, j = minWord.index1; j <= minWord.index2; i++, j++)
		newStr[i] = str[j];
	newStr[minWord.length] = '\0';

	cout << newStr;

	delete [] newStr;
	newStr = NULL;

	return 0;
}

bool isSymbol(const char* symb, int len, char C)
{
	for (int i = 0; i < len; i++)
		if (symb[i] == C)
			return true;

	return false;
}