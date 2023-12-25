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
	//Task("String61");
	char str[1000];
	cin.getline(str, 1000);

	int index = 0;
	Word word;
	word.index1 = word.index2 = -1;
	while (str[index] != '\0')
	{
		if (str[index] == '\\')
		{
			if (word.index2 < 0)
				word.index2 = index;
			else
			{
				word.index1 = word.index2;
				word.index2 = index;
			}
		}
		index++;
	}
	
	char* newStr = NULL;
	if (word.index1 < 0)
	{
		newStr = new char [2];
		newStr[0] = '\\';
		newStr[1] = '\0';
	}
	else
	{
		word.index1++;
		word.index2--;
		word.length = word.index2 - word.index1 + 1;
		
		newStr = new char [word.length + 1];
		for (int i = 0, j = word.index1; j <= word.index2; i++, j++)
			newStr[i] = str[j];
		newStr[word.length] = '\0';
	}

	cout << newStr;

	delete [] newStr;
	newStr = NULL;

	return 0;
}
