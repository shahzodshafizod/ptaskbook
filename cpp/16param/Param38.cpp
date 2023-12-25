#include <iostream>
using namespace std;

int PosLast(const char* S0, const char* S);

int main()
{
	//Task("Param38");
	const int stringSize = 100;
	char S0[stringSize+1], S[stringSize+1];
	for (int i = 1; i < 6; i++)
	{
		cin.getline(S0, stringSize+1);
		cin.getline(S, stringSize+1);
		cout << PosLast(S0, S);
	}
	
	return 0;
}

int PosLast(const char* S0, const char* S)
{
	int len = strlen(S), S0Len = strlen(S0), lastPos = 0;
	for (int i = 0; i < len; ++i)
	{
		if (S[i] == S0[0])
		{
			int j = i, index = 0;
			while ( (j < len) && (index < S0Len) && (S[j] == S0[index]) )
				index++, j++;
			if (index == S0Len)
			{
				lastPos = j-S0Len+1;
				i += S0Len;
				i--;
			}
		}
	}
	
	return lastPos;
}