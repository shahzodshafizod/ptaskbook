#include <iostream>
using namespace std;

int CharToInt(char C) { return int(C) - int('0'); }

int main()
{
	//Task("String23");
	char str[1000];
	cin.getline(str, 1000);

	int result = CharToInt(str[0]), index = 1;
	while (str[index] != '\0')
	{
		if (str[index] == '+')
			result += CharToInt(str[index+1]);
		else
			result -= CharToInt(str[index+1]);

		index += 2;
	}

	cout << result;
	
	return 0;
}