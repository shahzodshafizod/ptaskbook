#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void Solve()
{
    Task("Dynamic30");
	TNode* P1;
	pt >> P1;
	P1->Prev = NULL;
	TNode* temp;
	while ( (P1 != NULL) && (P1->Next != NULL) )
	{
		temp = P1;
		P1 = P1->Next;
		P1->Prev = temp;
		temp = NULL;
	}
	pt << P1;
}
