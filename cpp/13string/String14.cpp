#include <iostream>
using namespace std;

int main()
{
	//Task("String14");
	char str[1000];
	cin.getline(str, 1000);

	int index = 0, LETTERS = 0;
	while (str[index] != '\0')
	{
		LETTERS += (int)( (str[index] >= 'A') && (str[index] <= 'Z') );
		index++;
	}

	cout << LETTERS;
	
	return 0;
}
