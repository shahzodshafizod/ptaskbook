#include <iostream>
using namespace std;

int main()
{
	//Task("String18");
	char str[1000];
	cin.getline(str, 1000);

	int farq, index = 0;
	while (str[index] != '\0')
	{
		if ( (str[index] >= 'a') && (str[index] <= 'z') )
		{
			farq = int(str[index]) - int('a');
			str[index] = char( int('A') + farq );
		}
		else if ( (str[index] >= 'A') && (str[index] <= 'Z') )
		{
			farq = int(str[index]) - int('A');
			str[index] = char( int('a') + farq );
		}

		else if ( (str[index] >= 'а') && (str[index] <= 'я') )
		{
			farq = int(str[index]) - int('а');
			str[index] = char( int('А') + farq );
		}
		else if ( (str[index] >= 'А') && (str[index] <= 'Я') )
		{
			farq = int(str[index]) - int('А');
			str[index] = char( int('а') + farq );
		}
		else if (str[index] == 'ё')
			str[index] == 'Ё';
		else if (str[index] == 'Ё')
			str[index] == 'ё';

		index++;
	}

	cout << str;
	
	return 0;
}
