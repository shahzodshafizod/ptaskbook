#include <iostream>
using namespace std;

int main()
{
	//Task("String16");
	char str[1000];
	cin.getline(str, 1000);

	int index = 0;
	while (str[index] != '\0')
	{
		if ( (str[index] >= 'A') && (str[index] <= 'Z') )
		{
			int farq = int(str[index]) - int('A');
			str[index] = char( int('a') + farq );
		}
		index++;
	}

	cout << str;
	
	return 0;
}
