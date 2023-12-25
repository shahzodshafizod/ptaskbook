#include <iostream>
using namespace std;

int main()
{
	//Task("String53");
	char str[1000], mass[] = "!()-?.,\":;\'";
	int arraySize = sizeof(mass) / sizeof(mass[0]);
	cin.getline(str, 1000);

	int len = 0;
	while (str[len] != '\0')
		len++;

	int symbs = 0;
	for (int i = 0; i < len; i++)
	{
		bool isTrue = false;
		for (int j = 0; j < arraySize; j++)
			if (str[i] == mass[j])
			{
				isTrue = true;
				break;
			}
		symbs += (int)isTrue;
	}

	cout << symbs;
	
	return 0;
}
