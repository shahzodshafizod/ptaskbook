#include <iostream>
using namespace std;

int main()
{
	//Task("String13");
	char str[1000];
	cin.getline(str, 1000);

	int index = 0, digits = 0;
	while (str[index] != '\0')
	{
		digits += (int)((str[index] >= '0') && (str[index] <= '9'));
		index++;
	}

	cout << digits;
	
	return 0;
}
