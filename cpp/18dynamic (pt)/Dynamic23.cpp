#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void Solve()
{
    Task("Dynamic23");
	TNode *P1, *P2, *P3, *P4;
	pt >> P1 >> P2 >> P3 >> P4;
	while ( (P1 != NULL) && (P1->Data % 2 != 0) )
	{
		TNode* temp = P1;
		P1 = P1->Next;
		temp->Next = NULL;
		if (P3 == NULL)
			P3 = P4 = temp;
		else
		{
			P4->Next = temp;
			P4 = temp;
		}
		temp = NULL;
	}
	if (P1 == NULL)
		P2 = NULL;
	pt << P1 << P2 << P3 << P4;
}
