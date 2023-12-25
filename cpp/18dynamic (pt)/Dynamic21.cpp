#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void Solve()
{
    Task("Dynamic21");
	TNode *P1, *P2, *P3, *P4;
	pt >> P1 >> P2 >> P3 >> P4;
	while (P1 != NULL)
	{
		TNode* temp = P1;
		P1 = P1->Next;
		temp->Next = NULL;
		if (P4 == NULL)
			P3 = P4 = temp;
		else
		{
			P4->Next = temp;
			P4 = temp;
		}
		temp = NULL;
	}
	P2 = NULL;
	pt << P3 << P4;
}
