#include <iostream>
using namespace std;

struct Lifo
{
	char C;
	Lifo *Prev;
};

int GetErrorNo(const char* str, int& index, int len, Lifo** stack);
int CharToInt(char C);
void AddToStack(Lifo **stack, char C);
bool PopFromStack(Lifo **stack);

int main()
{
	//Task("Recur19");
	char str[100];
	cin >> str;
	int len = 0;
	while (str[len] != '\0')
		len++;
	Lifo *stack = NULL;
	int index = 0;
	int errorNo = GetErrorNo(str, index, len, &stack) + 1;
	if (stack != NULL)
	{
		if ( (errorNo == 0) || (index == len) )
			errorNo = len+1;
	}
	cout << errorNo;
	return 0;
}

int GetErrorNo(const char* str, int& index, int len, Lifo** stack)
{
	if (index >= len)
		return index;
	int errorNo = -1;
	if (str[index] == '(')
	{
		AddToStack(stack, str[index]);
		errorNo = GetErrorNo(str, ++index, len, stack);
	}
	else if (CharToInt(str[index]) < 0)
		errorNo = index;
	
	if ( (errorNo >= 0) || (index >= len-1) )
		return errorNo;

	switch (str[index+1])
	{
		case ')':
			if (PopFromStack(stack))
			{
				index++;
				if ( (index >= 2) && (str[index-2] == '(') )
					return index;
				return -1;
			}
			else
			{
				AddToStack(stack, str[index+1]);
				return index+1;
			}
			break;
		case '+': case '-': case '*':
			index += 2;
			break;
		default:
			return index+1;
	}
	return GetErrorNo(str, index, len, stack);
}

int CharToInt(char C)
{
	int digit = int(C) - int('0');
	if ( (digit >= 0) && (digit <= 9) )
		return digit;
	return -1;
}

void AddToStack(Lifo **stack, char C)
{
	Lifo *temp = new Lifo;
	temp->C = C;
	temp->Prev = *stack;
	*stack = temp;
	temp = NULL;
}

bool PopFromStack(Lifo **stack)
{
	if (*stack == NULL)
		return false;
	if ((*stack)->C == '(')
	{
		Lifo *temp = *stack;
		*stack = (*stack)->Prev;
		delete temp;
		return true;
	}
	return false;
}