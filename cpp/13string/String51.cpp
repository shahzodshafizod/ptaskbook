#include <iostream>
using namespace std;

struct Word
{
	int index1;
	int index2;
};

void sortWords(const char* str, Word mass[], int N);
int strCompare(const char* str, Word word1, Word word2);

int main()
{
	//Task("String51");
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

	// to fint the first and the last positions for every word
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
	newLen--; // there is no a space after the last word

	sortWords(str, mass, words);

	char* newStr = new char [newLen];
	index = 0;
	for (int i = 0; i < words; i++)
	{
		for (int j = mass[i].index1; j <= mass[i].index2; j++)
			newStr[index++] = str[j];
		if (i < words-1)
			newStr[index++] = ' ';
	}

	for (int i = 0; i < newLen; i++)
		str[i] = newStr[i];
	str[newLen] = '\0';

	delete [] newStr;
	newStr = NULL;

	cout << str;
	
	return 0;
}

void sortWords(const char* str, Word mass[], int N)
{
	for (int i = 0; i < N-1; i++)
		for (int j = 1; j < N-i; j++)
			if (strCompare(str, mass[j-1], mass[j]) == 1)
			{
				Word temp = mass[j-1];
				mass[j-1] = mass[j];
				mass[j] = temp;
			}
}

int strCompare(const char* str, Word word1, Word word2)
{
	bool afz = true;
	for (int i = word1.index1, j = word2.index1; (i <= word1.index2) && (j <= word2.index2); i++, j++)
	{
		if ( str[i] > str[j] )
			return 1;
		else if ( str[i] < str[j] )
			return -1;
	}

	int len1 = word1.index2 - word1.index1 + 1;
	int len2 = word2.index2 - word2.index1 + 1;
	
	if ( len1 > len2 )
		return 1;
	else if (len1 < len2)
		return -1;

	return 0;
}