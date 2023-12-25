#include <iostream>
using namespace std;

int main()
{
	//Task("String30");
	char C, S[1000], S0[1000];
	cin >> C;
	cin.ignore();
	cin.getline(S, 1000);
	cin.getline(S0, 1000);

	int len = 0, Cs = 0;
	while (S[len] != '\0')
	{
		if (S[len] == C)
			Cs++;
		len++;
	}

	int S0Len = 0;
	while (S0[S0Len] != '\0')
		S0Len++;

	int newLen = len+Cs*S0Len;
	char* newStr = new char [newLen+1];
	
	for (int i = 0, j = 0; i < len; i++)
	{
		newStr[j++] = S[i];
		if (S[i] == C)
			for (int k = 0; k < S0Len; k++)
				newStr[j++] = S0[k];
	}
	newStr[newLen] = '\0';

	cout << newStr;

	delete [] newStr;
	newStr = NULL;
	
	return 0;
}
