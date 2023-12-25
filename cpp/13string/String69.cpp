#include <iostream>
using namespace std;

struct Qavs
{
	char q;
	Qavs* prev;
};

void pushQavs(Qavs**, char);
void popQavs(Qavs**);
void Doiravi(Qavs**, char, int, int&);

int main()
{
	//Task("String69");
	char str[1000];
	cin.getline(str, 1000);

	int length = strlen(str);

	Qavs* qavs = NULL;
	int opens = 0, closes = 0, index = -2;
	for (int i = 0; i < length; i++)
	{
		if ( (str[i] == '(') || (str[i] == ')') )
			Doiravi(&qavs, str[i], i, index);		
	}

	index++;

	if ( ! qavs )
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

void Doiravi(Qavs** qavs, char C, int i, int& index)
{
	if (C == '(')
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
		if ( (*qavs) && ((*qavs)->q == '(') )
			popQavs(qavs);
		else
		{
			if (index < 0)
				index = i;

			pushQavs(qavs, C);
		}
	}
}