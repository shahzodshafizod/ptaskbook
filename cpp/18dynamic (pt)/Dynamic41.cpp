#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void Solve()
{
    Task("Dynamic41");
	TNode* P0;
	pt >> P0;
	TNode* prev = P0->Prev;
	TNode* next = P0->Next;
	if (prev != NULL)
		prev->Next = next;
	if (next != NULL)
		next->Prev = prev;
	delete P0;
	P0 = NULL;
	pt << prev << next;
}
