#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void Solve()
{
    Task("Dynamic18");
	int D;
	pt >> D;
	TNode *P1, *P2;
	pt >> P1 >> P2;
	TNode *temp = new TNode;
	temp->Data = D;
	temp->Next = NULL;
	P2->Next = temp;
	P2 = temp;
	temp = P1;
	P1 = P1->Next;
	D = temp->Data;
	delete temp;
	pt << D << P1 << P2;
}
