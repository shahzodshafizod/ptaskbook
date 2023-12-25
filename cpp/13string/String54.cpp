#include <iostream>
using namespace std;

int main()
{
	//Task("String54");
	char str[1000];
	cin.getline(str, 1000);
	int len = 0;
	while (str[len] != '\0')
		len++;

	//char mass[] = "!?.,():;-\"";
	char mass[] = "АаЕеЁёИиОоУуЭэЮюЯяЫы";
	int arraySize = sizeof(mass) / sizeof(mass[0]);
	
	int sadonoks = 0;
	for (int i = 0; i < len; )
	{
		while ( (i < len) && (str[i] == ' ') )
			i++;
		
		while ( (i < len) && (str[i] != ' ') )
		{
			for (int j = 0; j < arraySize; j++)
				if (str[i] == mass[j])
				{
					sadonoks++;
					break;
				}
			i++;
		}
	}

	cout << sadonoks;
	
	return 0;
}
