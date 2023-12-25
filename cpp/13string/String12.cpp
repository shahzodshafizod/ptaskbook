#include <iostream>
using namespace std;

int main()
{
	//Task("String12");
	char str[1000];
	cin.getline(str, 1000);
	int N;
	cin >> N;

	int len = 0;
	while (str[len] != '\0')
		len++;

	int newLen = len+N*(len-1);
	char* newStr = new char [newLen+1];
	for (int i = 0, j = 0; i < len; i++)
	{
		newStr[j++] = str[i];
		if (i < len-1)
			for (int k = 1; k <= N; k++)
				newStr[j++] = '*';
	}
	newStr[newLen] = '\0';

	cout << newStr;

	delete [] newStr;
	newStr = NULL;
	
	return 0;
}
