#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void Solve()
{
    Task("Dynamic34");
	int D;
	pt >> D;
	TNode* P0;
	pt >> P0;
	TNode* P = new TNode;
	P->Data = D;
	P->Prev = P0;
	P->Next = P0->Next;
	if (P0->Next != NULL)
		P0->Next->Prev = P;
	P0->Next = P;
	pt << P;
}
