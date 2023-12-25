#include <iostream>
using namespace std;

int main()
{
	//Task("String32");
	char S[1000], S0[1000];
	cin.getline(S, 1000);
	cin.getline(S0, 1000);

	int len = 0;
	while (S[len] != '\0')
		len++;

	int S0Len = 0;
	while (S0[S0Len] != '\0')
		S0Len++;

	int counts = 0;
	for (int i = 0; i < len; i++)
	{
		if (S[i] == S0[0])
		{
			int j = 0, k = i;
			bool ok = true;
			while ( (j < S0Len) && (k < len) )
				if (S[k++] != S0[j++])
				{
					ok = false;
					break;
				}
			if ( ok && (j == S0Len) )
				counts++;
		}
	}

	cout << counts;
	
	return 0;
}
