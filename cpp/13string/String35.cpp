#include <iostream>
using namespace std;

int main()
{
	//Task("String35");
	char S[1000], S0[1000];
	cin.getline(S, 1000);
	cin.getline(S0, 1000);

	int len = 0;
	while (S[len] != '\0')
		len++;

	int S0Len = 0;
	while (S0[S0Len] != '\0')
		S0Len++;

	for (int i = 0; i < len; i++)
	{
		if (S[i] == S0[0])
		{
			int j = 0, k = i;
			bool ok = true;
			while ( (j < S0Len) && (k < len) )
				if (S0[j++] != S[k++])
				{
					ok = false;
					break;
				}
			if ( ok && (j == S0Len) )
			{
				j = i;
				while (k < len)
					S[j++] = S[k++];
				len -= S0Len;
				S[len] = '\0';
				i--;
			}
		}
	}

	cout << S;
	
	return 0;
}
