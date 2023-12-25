#include <iostream>
using namespace std;

int main()
{
	//Task("String28");
	char C, S[1000];
	cin >> C;
	cin.ignore();
	cin.getline(S, 1000);

	int len = 0, Cs = 0;
	while (S[len] != '\0')
	{
		if (S[len] == C)
			Cs++;
		len++;
	}

	int newLen = len+Cs;
	char* newStr = new char [newLen+1];
	for (int i = 0, j = 0; i < len; i++)
	{
		newStr[j++] = S[i];
		if (S[i] == C)
			newStr[j++] = S[i];
	}
	newStr[newLen] = '\0';

	cout << newStr;

	delete [] newStr;
	newStr = NULL;
	
	return 0;
}
