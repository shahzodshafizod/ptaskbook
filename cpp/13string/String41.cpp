#include <iostream>
using namespace std;

int main()
{
	//Task("String41");
	char str[1000];
	cin.getline(str, 1000);

	int len = 0;
	while (str[len] != '\0')
		len++;

	int words = 0;
	for (int i = 0; i < len;)
	{
		while ( (i < len) && (str[i] == ' ') )
			i++;
		words++;
		while ( (i < len) && (str[i] != ' ') )
			i++;
	}

	cout << words;
	
	return 0;
}
