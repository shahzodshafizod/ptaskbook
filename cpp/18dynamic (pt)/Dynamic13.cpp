#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

struct TStack
{
	TNode* Top;
};

bool StackIsEmpty(TStack S);
int Peek(TStack S);
int Pop(TStack** S);

void Solve()
{
    Task("Dynamic13");
	TNode* P1;
	pt >> P1;
	TStack* stack = new TStack;
	stack->Top = P1;
	int counts = 1;
	while ( (counts < 6) && (!StackIsEmpty(*stack)) )
	{
		//pt << Peek(*stack);
		pt << Pop(&stack);
		counts++;
	}
	if (!StackIsEmpty(*stack))
	{
		pt << false;
		pt << stack->Top->Data << stack->Top;
	}
	else
		pt << true;
}

bool StackIsEmpty(TStack S)
{
	return (S.Top == NULL);
}

int Peek(TStack S)
{
	return S.Top->Data;
}

int Pop(TStack** S)
{
	TNode* temp = (*S)->Top;
	int D = temp->Data;
	(*S)->Top = (*S)->Top->Next;
	delete temp;
	return D;
}