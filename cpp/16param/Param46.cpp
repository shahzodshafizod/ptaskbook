#include <iostream>
using namespace std;

int BinToDec(const char* S);

int main()
{
	//Task("Param46");
	const int stringSize = 100;
	char str[stringSize+1];
	
	for (int i = 0; i < 5; i++)
	{
		cin.getline(str, stringSize+1);
		cout << BinToDec(str);
	}
	
	return 0;
}

int BinToDec(const char* S)
{
	int len = 0;
	while (S[len] != '\0')
		len++;
	
	int du = 1, result = 0;
	for (int j = len-1; j >= 0; j--)
	{
		if (S[j] == '1')
			result += du;
		du *= 2;
	}
	
	return result;
}