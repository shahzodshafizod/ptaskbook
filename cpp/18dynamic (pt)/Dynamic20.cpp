#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

int Pop(TNode** node);

void Solve()
{
    Task("Dynamic20");
	TNode *P1, *P2;
	pt >> P1 >> P2;
	while ( (P1 != NULL) && (P1->Data % 2 != 0) )
		pt << Pop(&P1);
	
	if (P1 == NULL)
		P2 = NULL;

	pt << P1 << P2;
}

int Pop(TNode** node)
{
	TNode* temp = *node;
	*node = (*node)->Next;
	int D = temp->Data;
	delete temp;
	return D;
}