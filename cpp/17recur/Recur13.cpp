#include <iostream>
using namespace std;

bool Palindrome(const char* S, int start, int end);

int main()
{
	//Task("Recur13");
	const int stringSize = 100;
	char str[stringSize+1];
	for (int i = 1; i < 6; i++)
	{
		cin.getline(str, stringSize+1);
		int len = 0;
		while (str[len] != '\0')
			len++;

		cout << Palindrome(str, 0, len-1);
	}
	return 0;
}

bool Palindrome(const char* S, int start, int end)
{
	if (start >= end)
		return true;

	bool areEqual = S[start] == S[end];
	return (areEqual ? Palindrome(S, start+1, end-1) : false );
}