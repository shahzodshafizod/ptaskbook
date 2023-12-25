#include <iostream>
using namespace std;

void SplitStr(const char* S, char*** W, int& N);

int main()
{
	//Task("Param41");
	const int stringSize = 100;
	char str[stringSize+1];
	cin.getline(str, stringSize+1);
	char** W = NULL;
	int N = 0;
	SplitStr(str, &W, N);
	
	cout << N;
	for (int i = 0; i < N; i++)
		cout << W[i];

	for (int i = 0; i < N; i++)
		delete [] W[i];
	delete [] W;
	W = NULL;
	
	return 0;
}

void SplitStr(const char* S, char*** W, int& N)
{
	int len = strlen(S);
	N = 0;
	for (int i = 0; i < len; i++)
	{
		while ( (i < len) && (S[i] == ' ') )
			i++;
		if (i == len)
			break;
		
		N++;
		while ( (i < len) && (S[i] != ' ') )
			i++;
	}
	(*W) = new char* [N];
	int from, to, wordLen, index = 0;
	for (int i = 0; i < len; i++)
	{
		while ( (i < len) && (S[i] == ' ') )
			i++;
		if (i == len)
			break;

		from = i;
		while ( (i < len) && (S[i] != ' ') )
			i++;
		to = i-1;

		wordLen = to - from + 1;
		(*W)[index] = new char [wordLen+1];
		for (int j = 0, k = from; k <= to; j++, k++)
			(*W)[index][j] = S[k];
		(*W)[index][wordLen] = '\0';
		index++;
	}
}