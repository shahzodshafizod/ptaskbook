#include <iostream>
using namespace std;

int main()
{
	//Task("String37");
	char S[1000], S1[1000], S2[1000];
	cin.getline(S, 1000);
	cin.getline(S1, 1000);
	cin.getline(S2, 1000);

	int len = 0;
	while (S[len] != '\0')
		len++;

	int S1Len = 0;
	while (S1[S1Len] != '\0')
		S1Len++;

	int index = -1;
	for (int i = 0; i < len; i++)
	{
		if (S[i] == S1[0])
		{
			int j = 0, k = i;
			bool ok = true;
			while ( (j < S1Len) && (k < len) )
				if (S1[j++] != S[k++])
				{
					ok = false;
					break;
				}
			if ( ok && (j == S1Len) )
			{
				index = k;
				i = k-1;
			}
		}
	}

	if (index >= 0)
	{
		int S2Len = 0;
		while (S2[S2Len] != '\0')
			S2Len++;
		
		int newLen = len - S1Len + S2Len;
		char* newStr = new char [newLen];
		
		int i;
		for (i = 0; i < index-S1Len; i++)
			newStr[i] = S[i];
		
		for (int j = 0; j < S2Len; j++)
			newStr[i++] = S2[j];
		
		for (int j = index; j < len; j++)
			newStr[i++] = S[j];

		for (int i = 0; i < newLen; i++)
			S[i] = newStr[i];
		S[newLen] = '\0';

		delete [] newStr;
		newStr = NULL;
	}

	cout << S;
	
	return 0;
}
