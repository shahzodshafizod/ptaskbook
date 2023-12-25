#include <iostream>
using namespace std;

void MoveRightK(char& C, int K);

int main()
{
	//Task("String63");
	char str[1000];
	cin.getline(str, 1000);

	int K;
	cin >> K;

	int index = 0;
	while (str[index] != '\0')
	{
		MoveRightK(str[index], K);
		index++;
	}

	cout << str;
	
	return 0;
}

void MoveRightK(char& C, int K)
{
	int code = int(C);
	
	if ( (C >= 'А') && (C <= 'Я') )
	{
		if (code+K > int('Я'))
			code = K - (int('Я') - code + 1) + int('А');
		else
			code += K;
	}
	else if ( (C >= 'а') && (C <= 'я') )
	{
		if (code+K > int('я'))
			code = K - (int('я') - code + 1) + int('а');
		else
			code += K;
	}
	
	C = char(code);
}