#include <iostream>
using namespace std;

int main()
{
	//Task("String68");
	char str[1000];
	cin.getline(str, 1000);

	int index = 0;
	bool initialized = false;
	char C;
	while (str[index] != '\0')
	{
		if ( (str[index] >= 'a') && (str[index] <= 'z') )
		{
			if (initialized == false)
			{
				C = str[index];
				initialized = true;
			}
			else if (C <= str[index])
				C = str[index];
			else
				break;
		}
		index++;
	}
	
	if (str[index] == '\0')
		cout << 0;
	else
		cout << (index+1);

	return 0;
}
