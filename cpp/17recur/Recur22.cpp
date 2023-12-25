#include <iostream>
using namespace std;

int GetMinMax(const char* S, int& index);
int CharToInt(char C);

int main()
{
	//Task("Recur22");
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
	do
	{
		index += 2;
		int b = GetMinMax(S, index);
		if (maximal)
			a = a > b ? a : b;
		else
			a = a < b ? a : b;
	}
	while (S[index+1] != ')');
	index++;
	return a;
}

int CharToInt(char C)
{
	int digit = int(C) - int('0');
	if ( (digit >= 0) && (digit <= 9) )
		return digit;
	return -1;
}