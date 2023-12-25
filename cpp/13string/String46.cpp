#include <iostream>
using namespace std;

int main()
{
	//Task("String46");
	char str[1000];
	cin.getline(str, 1000);

	int len = 0;
	while (str[len] != '\0')
		len++;

	int maxLength = -1;
	for (int i = 0; i < len; )
	{
		while ( (i < len) && (str[i] == ' ') )
			i++;
		int length = 0;
		while ( (i < len) && (str[i] != ' ') )
			length++, i++;

		if (maxLength < 0)
			maxLength = length;
		else if (length > maxLength)
			maxLength = length;
	}

	cout << maxLength;
	
	return 0;
}
