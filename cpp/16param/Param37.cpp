#include <iostream>
using namespace std;

int PosSub(const char* S0, const char* S, int K, int N);

int main()
{
	//Task("Param37");
	const int stringSize = 100;
	char S0[stringSize+1], S[stringSize+1];
	cin.getline(S0, stringSize+1);
	cin.getline(S, stringSize+1);

	int K, N;
	for (int i = 0; i < 3; i++)
	{
		cin >> K >> N;
		cout << PosSub(S0, S, K, N);
	}
	
	return 0;
}

int PosSub(const char* S0, const char* S, int K, int N)
{
	int from, to;
	int len = strlen(S);
	if (K > len)
	{
		from = to = 0;
	}
	else if (len - K < N)
	{
		from = K - 1;
		to = len - 1;
	}
	else
	{
		from = K - 1;
		to = from + N - 1;
	}

	int S0Len = strlen(S0);
	for (int i = from; i <= to; i++)
	{
		if (S[i] == S0[0])
		{
			int index = 0, j = i;
			while ( (j <= to) && (S0[index] == S[j]) )
				j++, index++;
			if (index == S0Len)
				return j-S0Len+1;
		}
	}

	return 0;
}