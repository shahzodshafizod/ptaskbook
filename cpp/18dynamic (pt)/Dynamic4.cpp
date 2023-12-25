#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void Solve()
{
	Task("Dynamic4");
	int N, number;
	pt >> N;
	pt >> number;
	TNode* stack = new TNode;
	stack->Data = number;
	stack->Next = NULL;
	for (int i = 2; i <= N; i++)
	{
		pt >> number;
		TNode* temp = new TNode;
		temp->Data = number;
		temp->Next = stack;
		stack = temp;
		temp = NULL;
	}
	pt << stack;
}
