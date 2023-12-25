#include <iostream>
using namespace std;

int main()
{
	//Task("String49");
	char str[1000];
	cin.getline(str, 1000);

	int len = 0;
	while (str[len] != '\0')
		len++;

	for (int i = len-1; i >= 0; i--)
	{
		while ( (i >= 0) && (str[i] == ' ') )
			i--;
		char C = str[i--];
		while ( (i >= 0) && (str[i] != ' ') )
		{
			if (str[i] == C)
				str[i] = '.';
			i--;
		}
	}

	cout << str;
	
	return 0;
}
