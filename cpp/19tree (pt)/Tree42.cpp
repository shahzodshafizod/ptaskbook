#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void DeleteNodes(TNode*, int);
void DeleteTree(TNode**);

void Solve()
{
    Task("Tree42");
	TNode* P1;
	pt >> P1;
	DeleteNodes(P1, P1->Data);
}

void DeleteNodes(TNode* P, int N)
{
	if (P == NULL)
		return;
	//================================================
	if ( (P->Left != NULL) && (P->Left->Data < N) )
		DeleteTree(&(P->Left));
	else
		DeleteNodes(P->Left, N);
	//================================================
	if ( (P->Right != NULL) && (P->Right->Data < N) )
		DeleteTree(&(P->Right));
	else
		DeleteNodes(P->Right, N);
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