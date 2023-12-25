#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void Solve()
{
    Task("Dynamic29");
	TNode* P2;
	pt >> P2;
	TNode *P1 = P2->Prev;
	TNode* P3 = P2->Next;
	pt << P1->Data << P3->Data;
	pt << P1 << P3;
}
