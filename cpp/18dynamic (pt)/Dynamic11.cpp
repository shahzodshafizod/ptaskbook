#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

struct TStack
{
	TNode* Top;
};

void Push(TStack** S, int D);

void Solve()
{
    Task("Dynamic11");
	TNode* P1;
	pt >> P1;
	int N, number;
	pt >> N;
	TStack* stack = new TStack;
	stack->Top = P1;
	for (int i = 1; i <= N; i++)
	{
		pt >> number;
		Push(&stack, number);
	}
	pt << stack->Top;
}

void Push(TStack** S, int D)
{
	TNode* newNode = new TNode;
	newNode->Data = D;
	newNode->Next = (*S)->Top;
	(*S)->Top = newNode;
	newNode = NULL;
}