#include <iostream>
using namespace std;

int main()
{
	//Task("String66");
	char str[1000];
	cin.getline(str, 1000);

	int len = 0;
	while (str[len] != '\0')
		len++;

	char* newStr = new char [len];
	int i = 0, j = 0;
	for (; i < len/2; i++)
	{
		newStr[len-1-i] = str[j++];
		newStr[i] = str[j++];
	}
	if (len % 2 != 0)
		newStr[i] = str[j];

	
	for (int i = 0; i < len; i++)
		str[i] = newStr[i];

	delete [] newStr;
	newStr = NULL;

	cout << str;
	
	return 0;
}
