#include <iostream>
using namespace std;

void TrimLeftC(char* S, char C);

int main()
{
	//Task("Param34");
	const int stringSize = 100;
	char str[stringSize+1];
	char C;
	cin >> C;
	cin.ignore();
	for (int i = 0; i < 5; i++)
	{
		cin.getline(str, stringSize+1);
		TrimLeftC(str, C);
		cout << str;
	}
	
	return 0;
}

void TrimLeftC(char* S, char C)
{
	if (C != S[0])
		return;
	int len = strlen(S);
	int index = 0;
	while ( (index < len) && (S[index] == C) )
		index++;
	
	int i = 0;
	for (int j = index; j < len; i++, j++)
		S[i] = S[j];
	S[i] = '\0';
}