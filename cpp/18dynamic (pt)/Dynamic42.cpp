#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void Solve()
{
    Task("Dynamic42");
	TNode* P1;
	pt >> P1;
	TNode *next, *i = P1;
	P1 = P1->Next;
	for (; i != NULL; i = i->Next)
	{
		if (i->Next == NULL)
		{
			if (i->Prev != NULL)
				i->Prev->Next = NULL;
			delete i;
			break;
		}
		next = i->Next;
		next->Prev = i->Prev;
		if (i->Prev != NULL)
			i->Prev->Next = next;
		delete i;
		i = next;
		next = NULL;
	}
	pt << P1;
}
