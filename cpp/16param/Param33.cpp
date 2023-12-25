#include <iostream>
using namespace std;

void LowCaseRus(char* S);

int main()
{
	//Task("Param33");
	const int stringSize = 100;
	char str[stringSize+1];
	for (int i = 0; i < 5; i++)
	{
		cin.getline(str, stringSize+1);
		LowCaseRus(str);
		cout << str;
	}
	
	return 0;
}

void LowCaseRus(char* S)
{
	int index = 0;
	while (S[index] != '\0')
	{
		if ( (S[index] >= 'А') && (S[index] <= 'Я') )
		{
			int farq = int(S[index]) - int('А');
			S[index] = char(int('а') + farq);
		}
		index++;
	}
}