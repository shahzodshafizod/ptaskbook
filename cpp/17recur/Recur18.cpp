#include <iostream>
using namespace std;

struct Fifo
{
	char C;
	int index;
	Fifo* Prev;
};

bool Validation(const char* str, int& index, int len, Fifo** stack);
int CharToInt(char C);
void AddToStack(Fifo** stack, char C);
bool PopFromStack(Fifo** stack, char C = '(');

int main()
{
	//Task("Recur18");
	char str[100];
	cin >> str;
	int len = strlen(str);
	Fifo* stack = new Fifo;
	stack->Prev = NULL;
	stack->index = 0;
	int index = 0;
	bool isValid = Validation(str, index, len, &stack);
	isValid = isValid && (stack->index == 0);
	cout << isValid;
	return 0;
}

bool Validation(const char* str, int& index, int len, Fifo** stack)
{
	bool isValid = str[index] == '(' ? AddToStack(stack, str[index]),
			Validation(str, ++index, len, stack)
			: CharToInt(str[index]) >= 0;
	
	if (index >= len-1) return isValid;

	switch (str[index+1])
	{
		case ')': index += (int)PopFromStack(stack); return isValid;
		case '+': case '-': case '*': index += 2; break;
		default: isValid = false;
	}
	return isValid && Validation(str, index, len, stack);
}

int CharToInt(char C)
{
	int digit = int(C) - int('0');
	if ( (digit >= 0) && (digit <= 9) )
		return digit;
	return -1;
}

void AddToStack(Fifo** stack, char C)
{
	Fifo* temp = new Fifo;
	temp->C = C;
	temp->index = (*stack)->index + 1;
	temp->Prev = *stack;
	*stack = temp;
}

bool PopFromStack(Fifo** stack, char C)
{
	if ( ((*stack)->index != 0) && ((*stack)->C == C) )
	{
		Fifo* temp = (*stack)->Prev;
		delete (*stack);
		(*stack) = temp;
		return true;
	}
	AddToStack(stack, ')');
	return false;
}