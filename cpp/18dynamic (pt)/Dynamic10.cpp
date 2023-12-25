#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void Solve()
{
    Task("Dynamic10");
	TNode *P1, *temp;
	pt >> P1;
	TNode *evens = NULL, *unevens = NULL;
	while (P1 != NULL)
	{
		temp = P1;
		P1 = P1->Next;
		if (temp->Data % 2 == 0)
		{
			temp->Next = evens;
			evens = temp;
		}
		else
		{
			temp->Next = unevens;
			unevens = temp;
		}
		temp = NULL;
	}
	pt << evens << unevens;
}
