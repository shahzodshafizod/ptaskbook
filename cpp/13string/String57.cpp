#include <iostream>
using namespace std;

int main()
{
	//Task("String57");
	char str[1000];
	cin.getline(str, 1000);

	int len = 0;
	while (str[len] != '\0')
		len++;

	for (int i = 0; i < len; )
	{
		if ( (i < len) && str[i] == ' ' )
			i++;

		while ( (i < len) && (str[i] == ' ') )
		{
			for (int j = i; j < len-1; j++)
				str[j] = str[j+1];
			len--;
			str[len] = '\0';
		}

		while ( (i < len) && (str[i] != ' ') )
			i++;
	}

	cout << str;
	
	return 0;
}
