#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void Solve()
{
    Task("Dynamic72");
	TNode* P1, *P2;
	pt >> P1 >> P2;
	TNode* P3 = P1->Prev;
	TNode* P4 = P2->Prev;
	if (P2 != P4)
	{
		TNode* barer2 = P2;
		P2 = P2->Next;
		if (P2 != NULL)
		{
			P4->Next = P1;
			P1->Prev = P4;
			P3->Next = P2;
			P2->Prev = P3;
			P2 = P4;
		}
		delete barer2;
	}
	else
	{
		delete P2;
		P2 = P1->Prev;
	}
	P1 = P1->Next;
	pt << P1 << P2;
}
