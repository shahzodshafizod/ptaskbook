#include <iostream>
using namespace std;

int main()
{
	//Task("String43");
	char str[1000];
	cin.getline(str, 1000);

	int len = 0;
	while (str[len] != '\0')
		len++;

	int words = 0;
	for (int i = 0; i < len; )
	{
		while ( (i < len) && (str[i] == ' ') )
			i++;
		bool hasA = false;
		while ( (i < len) && (str[i] != ' ') )
		{
			if (str[i] == 'А')
				hasA = true;
			i++;
		}

		words += (int)hasA;
	}

	cout << words;
	
	return 0;
}
