#include <iostream>
using namespace std;

int main()
{
	//Task("String45");
	char str[1000];
	cin.getline(str, 1000);

	int len = 0;
	while (str[len] != '\0')
		len++;

	int length, minLength = -1;
	for (int i = 0; i < len; )
	{
		while ( (i < len) && (str[i] == ' ') )
			i++;
		length = 0;
		while ( (i < len) && (str[i] != ' ') )
			i++, length++;
		
		if (minLength < 0)
			minLength = length;
		else if (length < minLength)
			minLength = length;
	}

	cout << minLength;
	
	return 0;
}
