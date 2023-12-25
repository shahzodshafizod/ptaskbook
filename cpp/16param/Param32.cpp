#include <iostream>
using namespace std;

void UpCaseRus(char* S);

int main()
{
	//Task("Param32");
	const int stringSize = 100;
	char str[stringSize+1];
	for (int i = 1; i < 6; i++)
	{
		cin.getline(str, stringSize+1);
		UpCaseRus(str);
		cout << str;
	}
	
	return 0;
}

void UpCaseRus(char* S)
{
	int index = 0;
	while (S[index] != '\0')
	{
		if ( (S[index] >= 'а') && (S[index] <= 'я') )
		{
			int farq = int(S[index]) - int('а');
			S[index] = char(int('А') + farq);
		}
		index++;
	}
}