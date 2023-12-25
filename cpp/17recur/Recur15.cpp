#include <iostream>
using namespace std;

int Rec(const char* S, int& index);
int Zarb(const char* S, int& index);
int CharToInt(char C);

int main()
{
	//Task("Recur15");
	const int stringSize = 100;
	char str[stringSize+1];
	cin >> str;
	int index = strlen(str) - 1;
	int result = Rec(str, index);
	cout << result;
	return 0;
}

int Rec(const char* S, int& index)
{
	int number = (index != 0) && (S[index-1] != '*') ? CharToInt(S[index]) : Zarb(S, index);
	
	if (index == 0)
		return number;
	
	index -= 2;
	
	if (S[index+1] == '+')
		return Rec(S, index) + number;
	
	else
		return Rec(S, index) - number;
}

int Zarb(const char* S, int& index)
{
	int number = CharToInt(S[index]);
	
	if ( (index == 0) || (S[index-1] != '*') )
		return number;
	
	index -= 2;
	
	return Zarb(S, index) * number;
}

int CharToInt(char C)
{
	int farq = int(C) - int('0');
	if ( (farq >= 0) && (farq <= 9) )
		return farq;
	return -1;
}