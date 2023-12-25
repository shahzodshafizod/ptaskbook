#include <iostream>
using namespace std;

int main()
{
	//Task("String17");
	char str[1000];
	cin.getline(str, 1000);

	int index = 0;
	while (str[index] != '\0')
	{
		if ( (str[index] >= 'a') && (str[index] <= 'z') )
		{
			int farq = int(str[index]) - int('a');
			str[index] = char( int('A') + farq );
		}
		else if ( (str[index] >= 'а') && (str[index] <= 'я') )
		{
			int farq = int(str[index]) - int('а');
			str[index] = char( int('А') +farq );
		}
		
		index++;
	}

	cout << str;
	
	return 0;
}
