#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void Solve()
{
    Task("Dynamic9");
	TNode *P1, *P2, *temp;
	pt >> P1 >> P2;
	while ( (P1 != NULL) && (P1->Data % 2 != 0) )
	{
		temp = P1;
		P1 = P1->Next;
		temp->Next = P2;
		P2 = temp;
		temp = NULL;
	}
	pt << P1 << P2;
}
