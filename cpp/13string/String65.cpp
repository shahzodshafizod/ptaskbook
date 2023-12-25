#include <iostream>
using namespace std;

void MoveRightK(char&, int);
void MoveLeftK(char&, int);

typedef void (*FuncPointer)(char&, int);

int main()
{
	//Task("String65");
	char str[1000], C;
	cin.getline(str, 1000);
	cin >> C;

	int index = 0;
	int K = str[0] - C;
	//void (*funcPtr)(char&, int) = MoveLeftK;
	FuncPointer funcPtr = MoveLeftK;
	if (K < 0)
	{
		funcPtr = MoveRightK;
		K *= -1;
	}
	while (str[index] != '\0')
	{
		funcPtr(str[index], K);
		index++;
	}

	cout << K;
	cout << str;
	
	return 0;
}

void MoveRightK(char& C, int K)
{
	int code = int(C);

	if ( (C >= 'А') && (C <= 'Я') )
	{
		if ( code + K > int('Я') )
			code = K - (int('Я') - code + 1) + int('А');
		else
			code += K;
	}
	else if ( (C >= 'а') && (C <= 'я') )
	{
		if ( code + K > int('я') )
			code = K - (int('я') - code + 1) + int('а');
		else
			code += K;
	}

	C = char(code);
}

void MoveLeftK(char& C, int K)
{
	int code = int(C);

	if ( (C >= 'А') && (C <= 'Я') )
	{
		if ( code - K < int('А') )
			code = int('Я') - (K - (code - int('А') + 1));
		else
			code -= K;
	}
	else if ( (C >= 'а') && (C <= 'я') )
	{
		if ( code - K < int('а') )
			code = int('я') - (K - (code - int('а') + 1));
		else
			code -= K;
	}

	C = char(code);
}