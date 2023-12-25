#include <iostream>
using namespace std;

void TrimRightC(char* S, char C);

int main()
{
	//Task("Param35");
	const int stringSize = 100;
	char str[stringSize+1];
	char C;
	cin >> C;
	cin.ignore();
	for (int i = 0; i < 5; i++)
	{
		cin.getline(str, stringSize+1);
		TrimRightC(str, C);
		cout << str;
	}
	
	return 0;
}

void TrimRightC(char* S, char C)
{
	int index = strlen(S)-1;
	if (S[index] != C)
		return;
	
	while ( (index >= 0) && (S[index] == C) )
		index--;
	
	S[++index] = '\0';
}