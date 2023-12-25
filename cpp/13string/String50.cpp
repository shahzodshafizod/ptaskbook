#include <iostream>
using namespace std;

struct Word
{
	int index1;
	int index2;
};

int main()
{
	//Task("String50");
	char str[1000];
	cin.getline(str, 1000);

	int len = 0;
	while (str[len] != '\0')
		len++;

	// to find a count of words
	int words = 0;
	for (int i = 0; i < len; )
	{
		while ( (i < len) && (str[i] == ' ') )
			i++;
		if (i == len)
			break;
		words++;
		while ( (i < len) && (str[i] != ' ') )
			i++;
	}
	
	// to find the first and the last positions for every word
	Word* mass = new Word [words];
	int index = 0, newLen = 0;
	for (int i = 0; i < len; )
	{
		while ( (i < len) && (str[i] == ' ') )
			i++;
		if (i == len)
			break;
		mass[index].index1 = i;
		while ( (i < len) && (str[i] != ' ') )
			i++;
		mass[index].index2 = i-1;
		
		newLen += mass[index].index2 - mass[index].index1 + 1;
		newLen++; // for a space after every word
		index++;
	}
	newLen--; // there is no space after the last word
	
	char* newStr = new char [newLen];
	index = 0;
	// to organize new string - buffer
	for (int i = words-1; i >= 0; i--)
	{
		for (int j = mass[i].index1; j <= mass[i].index2; j++)
			newStr[index++] = str[j];
		if (i > 0)
			newStr[index++] = ' ';
	}
	
	// to copy the new string to our string - buffer
	for (int i = 0; i < newLen; i++)
		str[i] = newStr[i];
	str[newLen] = '\0';

	// deleting the temperate string - buffer
	delete [] newStr;
	newStr = NULL;

	cout << str;

	return 0;
}
