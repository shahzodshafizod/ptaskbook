#include <iostream>
using namespace std;

int main()
{
	//Task("String59");
	char str[1000];
	cin.getline(str, 1000);

	int len = 0;
	while (str[len] != '\0')
		len++;

	int i = len-1;
	int length = 0;
	while ( (i >= 0) && (str[i] != '\\') && (str[i] != '.') )
	{
		i--;
		length++;
	}

	char* newStr = NULL;
	if (str[i] != '.')
	{
		newStr = new char [1];
		newStr[0] = '\0';
	}
	else
	{
		newStr = new char [length+1];
		i++;
		for (int j = 0; i < len; i++, j++)
			newStr[j] = str[i];
		newStr[length] = '\0';
	}

	cout << newStr;

	delete [] newStr;
	newStr = NULL;
	
	return 0;
}
