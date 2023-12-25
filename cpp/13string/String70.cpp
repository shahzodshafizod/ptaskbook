#include <iostream>
using namespace std;

struct Qavs
{
	char q;
	Qavs* prev;
};

void pushQavs(Qavs**, char);
void popQavs(Qavs**);
void Qavsho(Qavs**, char, int, int&, int);

int main()
{
	//Task("String70");
	char str[1000];
	cin.getline(str, 1000);

	int length = strlen(str);

	Qavs* qavsD = NULL;
	Qavs* qavsF = NULL;
	Qavs* qavsR = NULL;
	int opens = 0, closes = 0, index = -2;
	for (int i = 0; i < length; i++)
	{
		if ( (str[i] == '(') || (str[i] == ')') )
			Qavsho(&qavsD, str[i], i, index, 1);

		else if ( (str[i] == '{') || (str[i] == '}') )
			Qavsho(&qavsF, str[i], i, index, 2);

		else if ( (str[i] == '[') || (str[i] == ']') )
			Qavsho(&qavsR, str[i], i, index, 3);
	}

	index++;

	if ( !qavsD && !qavsF && !qavsR )
		cout << 0;
	else
		cout << index;
	
	return 0;
}

void popQavs(Qavs** qavs)
{
	Qavs* prevQ = (*qavs)->prev;
	
	delete (*qavs);
	
	(*qavs) = prevQ;
}

void pushQavs(Qavs** qavs, char C)
{
	Qavs* newQ = new Qavs;
	newQ->q = C;
	newQ->prev = (*qavs);
	
	(*qavs) = newQ;
}

void Qavsho(Qavs** qavs, char C, int i, int& index, int whatQavs)
{
	char q = whatQavs == 1 ? '(' : whatQavs == 2 ? '{' : '[';
	
	if (C == q)
	{
		if (! (*qavs) )
		{
			(*qavs) = new Qavs;
			(*qavs)->q = C;
			(*qavs)->prev = NULL;
		}
		else
			pushQavs(qavs, C);
	}
	else
	{
		if ( (*qavs) && ((*qavs)->q == q) )
			popQavs(qavs);
		else
		{
			if (index < 0)
				index = i;

			pushQavs(qavs, C);
		}
	}
}