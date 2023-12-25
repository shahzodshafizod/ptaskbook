#include <iostream>
using namespace std;

char* InvStr(const char* S, int K, int N);

int main()
{
	//Task("Param36");
	const int stringSize = 100;
	char S[stringSize+1];
	cin.getline(S, stringSize+1);
	int K, N;
	char* invedStr = NULL;
	for (int i = 1; i <= 3; i++)
	{
		cin >> K >> N;
		invedStr = InvStr(S, K, N);
		cout << invedStr;
		
		delete [] invedStr;
		invedStr = NULL;
	}
	
	return 0;
}

char* InvStr(const char* S, int K, int N)
{
	int newLen, from, to;
	int len = strlen(S);
	
	if (K > len)
		from = to = newLen = 0;
	else if (len-K < N)
	{
		from = K-1;
		to = len-1;
		newLen = to - from + 1;
	}
	else
	{
		from = K-1;
		to = from+N-1;
		newLen = to - from + 1;
	}
	
	char* newS = new char [newLen+1];
	for (int i = 0, j = to; j >= from; j--, i++)
		newS[i] = S[j];
	
	newS[newLen] = '\0';
	return newS;
}