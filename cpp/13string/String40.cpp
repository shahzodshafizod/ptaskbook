#include <iostream>
using namespace std;

int main()
{
	//Task("String40");
	char str[1000];
	cin.getline(str, 1000);

	int len = 0;
	while (str[len] != '\0')
		len++;

	int index1 = -1, index2 = -1;
	for (int i = 0; i < len; i++)
	{
		if (str[i] == ' ')
			if (index1 < 0)
				index1 = i;
			else
				index2 = i;
	}

	int newLen = 0;
	char* newStr = new char [newLen+1];
	
	if ( (index1 >= 0) && (index2 >= 0) )
	{
		newLen = index2 - index1 - 1;
		delete [] newStr;
		newStr = new char [newLen+1];
		for (int i = 0, j = index1+1; j < index2; i++, j++)
			newStr[i] = str[j];
	}
	newStr[newLen] = '\0';

	cout << newStr;

	delete [] newStr;
	newStr = NULL;
	
	return 0;
}
