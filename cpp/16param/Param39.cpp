#include <iostream>
using namespace std;

int PosK(const char* S0, const char* S, int K);

int main()
{
	//Task("Param39");
	const int stringSize = 100;
	char S0[stringSize+1], S[stringSize+1];
	int K;
	for (int i = 1; i < 6; i++)
	{
		cin.getline(S0, stringSize+1);
		cin.getline(S, stringSize+1);
		cin >> K;
		cin.ignore();
		cout << PosK(S0, S, K);
	}
	
	return 0;
}

int PosK(const char* S0, const char* S, int K)
{
	int len = strlen(S);
	int S0Len = strlen(S0);
	int counts = 0;

	for (int i = 0; i < len; i++)
	{
		if (S[i] == S0[0])
		{
			int j = i, index = 0;
			while ( (j < len) && (index < S0Len) && (S0[index] == S[j]) )
				j++, index++;
			if (index == S0Len)
			{
				counts++;
				if (counts == K)
					return i+1;
				i = j-1;
			}
		}
	}
	return 0;
}