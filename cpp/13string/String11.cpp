#include <iostream>
using namespace std;

int main()
{
	//Task("String11");
	char str[1000];
	cin.getline(str, 1000);

	int len = 0;
	while (str[len] != '\0')
		len++;

	char* newStr = new char [2*len];
	for (int i = 0, j = 0; i < len; i++)
	{
		newStr[j++] = str[i];
		if (i < len-1)
			newStr[j++] = ' ';
	}
	newStr[2*len-1] = '\0';

	cout << newStr;

	delete [] newStr;
	newStr = NULL;
	
	return 0;
}
