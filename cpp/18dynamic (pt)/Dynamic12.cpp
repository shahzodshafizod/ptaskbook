#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

struct TStack
{
	TNode* Top;
};

int Pop(TStack** S);

void Solve()
{
    Task("Dynamic12");
	TNode* P1;
	pt >> P1;
	TStack* stack = new TStack;
	stack->Top = P1;
	for (int i = 1; i < 6; i++)
		pt << Pop(&stack);
	
	pt << stack->Top;
}

int Pop(TStack** S)
{
	TNode* temp = (*S)->Top;
	int D = temp->Data;
	(*S)->Top = (*S)->Top->Next;
	delete temp;
	return D;
}