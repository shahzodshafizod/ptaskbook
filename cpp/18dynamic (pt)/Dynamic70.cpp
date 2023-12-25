#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void Solve()
{
    Task("Dynamic70");
	TNode* P1, *P2;
	pt >> P1 >> P2;
	TNode* barer = new TNode;
	barer->Data = 0;
	if (P1 != NULL)
	{
		barer->Next = P1;
		P1->Prev = barer;
	}
	else
		barer->Next = barer;
	
	if (P2 != NULL)
	{
		barer->Prev = P2;
		P2->Next = barer;
	}
	else
		barer->Prev = barer;

	pt << barer;
}
