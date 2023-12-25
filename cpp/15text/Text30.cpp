#include <iostream>
#include <fstream>
using namespace std;

void makeStr(const char* str, char** newStr, int length, int index1, int index2);
int main()
{
	//Task("Text30");
	const int stringSize = 80;
	char fileName[100], str[stringSize];
	cin.getline(fileName, 100);
	ifstream file(fileName);
	int strLen, index1, index2, length, minLength = -1;
	char* newStr = NULL;
	while (file.peek() != -1)
	{
		file.getline(str, stringSize);
		strLen = strlen(str);
		for (int i = 0; i < strLen;)
		{
			while ( (i < strLen) && (str[i] == ' ') )
				i++;
			length = 0;
			index1 = i;
			while ( (i < strLen) && (str[i] != ' ') )
			{
				length++;
				i++;
			}
			index2 = i-1;

			if (minLength < 0)
			{
				minLength = length;
				makeStr(str, &newStr, length, index1, index2);
			}
			else if (length <= minLength)
			{
				minLength = length;
				makeStr(str, &newStr, length, index1, index2);
			}
		}
	}
	file.close();
	cout << newStr;

	delete [] newStr;
	newStr = NULL;
	
	return 0;
}

void makeStr(const char* str, char** newStr, int length, int index1, int index2)
{
	delete [] (*newStr);
	(*newStr) = new char [length+1];
	for (int i = index1, k = 0; i <= index2; i++, k++)
		(*newStr)[k] = str[i];
	(*newStr)[length] = '\0';
}