#include <iostream>
using namespace std;

struct Word
{
	int index1;
	int index2;
	int length;
};

int main()
{
	//Task("String58");
	char str[1000];
	cin.getline(str, 1000);

	int len = 0;
	Word word;
	while (str[len] != '\0')
	{
		if (str[len] == '\\')
			word.index1 = len;
		len++;
	}

	word.index1++;
	word.index2 = word.index1;
	word.length = 0;
	while ( (word.index2 < len) && (str[word.index2] != '.') )
	{
		word.index2++;
		word.length++;
	}
	word.index2--;

	char* newStr = new char [word.length + 1];
	for (int i = 0, j = word.index1; j <= word.index2; i++, j++)
		newStr[i] = str[j];
	newStr[word.length] = '\0';

	cout << newStr;

	delete [] newStr;
	newStr = NULL;

	return 0;
}
