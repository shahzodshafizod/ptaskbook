#include <iostream>
using namespace std;

int main()
{
	//Task("String48");
	char str[1000];
	cin.getline(str, 1000);

	int len = 0;
	while (str[len] != '\0')
		len++;

	for (int i = 0; i < len; )
	{
		while ( (i < len) && (str[i] == ' ') )
			i++;
		char C = str[i++];
		while ( (i < len) && (str[i] != ' ') )
		{
			if (str[i] == C)
				str[i] = '.';
			i++;
		}
	}

	cout << str;
	
	return 0;
}
