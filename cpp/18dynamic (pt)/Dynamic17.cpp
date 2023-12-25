#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void Solve()
{
    Task("Dynamic17");
	int D;
	pt >> D;
	TNode *P1, *P2;
	pt >> P1 >> P2;
	TNode *temp = new TNode;
	temp->Data = D;
	temp->Next = NULL;
	
	if (P1 == NULL)
		P1 = temp;
	
	if (P2 == NULL)
		P2 = temp;
	else
	{
		P2->Next = temp;
		P2 = temp;
	}

	pt << P1 << P2;
}
