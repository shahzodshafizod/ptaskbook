#include <iostream>
using namespace std;

int main()
{
	//Task("String42");
	char str[1000];
	cin.getline(str, 1000);
	
	int len = 0;
	while (str[len] != '\0')
		len++;

	int words = 0;
	char firstChar, lastChar;
	for (int i = 0; i < len;)
	{
		while ( (i < len) && (str[i] == ' ') )
			i++;
		if (i == len)
			break;
		firstChar = str[i];
		while ( (i < len) && (str[i] != ' ') )
			i++;
		lastChar = str[i-1];
		if (firstChar == lastChar)
			words++;
	}

	cout << words;
	
	return 0;
}
