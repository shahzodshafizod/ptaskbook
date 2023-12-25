#include <iostream>
using namespace std;

int Rec(const char* S, int index);
int CharToInt(char C);

int main()
{
	//Task("Recur14");
	const int stringSize = 100;
	char str[stringSize+1];
	cin >> str;
	int len = 0;
	while (str[len] != '\0')
		len++;
	cout << Rec(str, len-1);
	return 0;
}

int Rec(const char* S, int index)
{
	int digit = CharToInt(S[index]);
	if (index == 0)
		return digit;
	
	if (S[index-1] == '+')
		return Rec(S, index-2) + digit;
	
	else
		return Rec(S, index-2) - digit;
}

int CharToInt(char C)
{
	int farq = int(C) - int('0');
	if ( (farq >= 0) && (farq <= 9) )
		return farq;
	return -1;
}