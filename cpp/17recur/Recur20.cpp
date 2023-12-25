#include <iostream>
using namespace std;

int GetMinMax(const char* S, int& index);
int CharToInt(char C);

int main()
{
	//Task("Recur20");
	char str[100];
	cin >> str;
	int index = 0;
	int minmax = GetMinMax(str, index);
	cout << minmax;
	return 0;
}

int GetMinMax(const char* S, int& index)
{
	int a = CharToInt(S[index]);
	if (a >= 0)
		return a;
	bool maximal = S[index] == 'M';
	index += 2;
	a = GetMinMax(S, index);
	index += 2;
	int b = GetMinMax(S, index);
	index++;
	
	return ( maximal ? a > b ? a : b : a < b ? a : b );
}

int CharToInt(char C)
{
	int digit = int(C) - int('0');
	if ( (digit >= 0) && (digit <= 9) )
		return digit;
	return -1;
}