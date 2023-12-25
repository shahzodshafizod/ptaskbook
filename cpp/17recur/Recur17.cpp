#include <iostream>
using namespace std;

int Rec(const char* S, int& index);
int Zarb(const char* S, int& index);
int CharToInt(char C);

int main()
{
	//Task("Recur17");
	const int stringSize = 100;
	char str[stringSize+1];
	cin >> str;
	int index = strlen(str)-1;
	int result = Rec(str, index);
	cout << result;
	return 0;
}

int Rec(const char* S, int& index)
{
	int number = S[index] == ')' ? Rec(S, --index) : CharToInt(S[index]);
	
	if (S[index-1] == '*')
	{
		index -= 2;
		number *= Zarb(S, index);
	}
	
	if (index == 0)
		return number;

	if (S[index-1] == '(')
	{
		index--;
		return number;
	}

	index -= 2;
	
	if (S[index+1] == '+')
		return Rec(S, index) + number;
	
	return Rec(S, index) - number;
}

int Zarb(const char* S, int& index)
{
	int number = S[index] == ')' ? Rec(S, --index) : CharToInt(S[index]);
	if ( (index == 0) || (S[index-1] != '*') )
		return number;
	
	index -= 2;
	return Zarb(S, index) * number;
}

int CharToInt(char C)
{
	int digit = int(C) - int('0');
	if ( (digit >= 0) && (digit <= 9) )
		return digit;
	return -1;
}