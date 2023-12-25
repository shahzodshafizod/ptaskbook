#include <iostream>
using namespace std;

void MoveLeftK(char& C, int K);

int main()
{
	//Task("String64");
	char str[1000];
	cin.getline(str, 1000);

	int K;
	cin >> K;

	int index = 0;
	while (str[index] != '\0')
	{
		MoveLeftK(str[index], K);
		index++;
	}

	cout << str;
	
	return 0;
}

void MoveLeftK(char& C, int K)
{
	int code = int(C);

	if ( (C >= 'а') && (C <= 'я') )
	{
		if (code-K < int('а'))
			code = int('я') - (K - (code - int('а') + 1));
		else
			code -= K;
	}
	else if ( (C >= 'А') && (C <= 'Я') )
	{
		if (code-K < int('А'))
			code = int('Я') - (K - (code - int('А') + 1));
		else
			code -= K;
	}

	C = char(code);
}