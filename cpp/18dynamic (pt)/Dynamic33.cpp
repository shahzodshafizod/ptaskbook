#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void Solve()
{
    Task("Dynamic33");
	int D;
	pt >> D;
	TNode* P0;
	pt >> P0;
	TNode* P = new TNode;
	P->Data = D;
	P->Next = P0;
	P->Prev = P0->Prev;
	if (P0->Prev != NULL)
		P0->Prev->Next = P;
	
	P0->Prev = P;
	
	pt << P;
}
