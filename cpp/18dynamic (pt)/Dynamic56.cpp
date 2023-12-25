#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void Solve()
{
    Task("Dynamic56");
	TNode* P1, *P2;
	pt >> P1 >> P2;
	TNode* P3 = P1, *P4 = P2;
	while (P3->Next != P4)
	{
		P3 = P3->Next;
		P4 = P4->Prev;
	}
	P3->Next = P1;
	P1->Prev = P3;

	P4->Prev = P2;
	P2->Next = P4;

	pt << P3 << P4;
}
