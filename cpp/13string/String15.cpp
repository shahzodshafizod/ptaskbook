#include <iostream>
using namespace std;

int main()
{
	//Task("String15");
	char str[1000];
	cin.getline(str, 1000);

	int index = 0, letters = 0;
	while (str[index] != '\0')
	{
		letters += (int)( (str[index] >= 'a') && (str[index] <= 'z') || (str[index] >= 'а') && (str[index] <= 'я') || (str[index] == 'ё') );
		index++;
	}

	cout << letters;
	
	return 0;
}
