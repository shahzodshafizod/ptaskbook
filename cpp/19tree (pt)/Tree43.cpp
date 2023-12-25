#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void DeleteNodes(TNode*);
void DeleteTree(TNode**);

void Solve()
{
    Task("Tree43");
	TNode* P1;
	pt >> P1;
	DeleteNodes(P1);
}

void DeleteNodes(TNode* P)
{
	if (P == NULL)
		return;
	//agar du shoxa doshta boshad,
	if ( (P->Left != NULL) && (P->Right != NULL) )
	{
		if (P->Data % 2 == 0)
			DeleteTree(&(P->Right));
		else
			DeleteTree(&(P->Left));
	}
	DeleteNodes(P->Left);
	DeleteNodes(P->Right);
}

void DeleteTree(TNode** P)
{
	if (*P == NULL)
		return;
	TNode* node = (*P)->Left;
	DeleteTree(&node);
	node = (*P)->Right;
	DeleteTree(&node);
	delete *P;
	*P = NULL;
}