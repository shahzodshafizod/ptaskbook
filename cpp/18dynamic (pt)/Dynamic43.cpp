#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void Solve()
{
    Task("Dynamic43");
	TNode* P1;
	pt >> P1;
	for (TNode* i = P1; i != NULL; i = i->Next)
	{
		if (i->Next == NULL)
		{
			if (i->Data % 2 != 0)
			{
				if (i->Prev != NULL)
					i->Prev->Next = i->Next;
				if (i->Next != NULL)
					i->Next->Prev = i->Prev;
				if (i == P1)
					P1 = P1->Next;
				delete i;
			}
			break;
		}
		if (i->Data % 2 != 0)
		{
			if (i->Prev != NULL)
				i->Prev->Next = i->Next;
			if (i->Next != NULL)
				i->Next->Prev = i->Prev;
			if (i == P1)
				P1 = P1->Next;
			delete i;
		}
	}
	pt << P1;
}
