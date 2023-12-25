#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void PutNodeBefore(TNode* newNode, TNode* node);

void Solve()
{
    Task("Dynamic51");
	TNode* P1, *P2, *P0;
	pt >> P1 >> P2 >> P0;
	
	while (P1 != NULL)
	{
		TNode* next = P1->Next;
		PutNodeBefore(P1, P0);
		P1 = next;
	}
	P2 = P0;
	while (P2->Next != NULL)
		P2 = P2->Next;
	P1 = P0;
	while (P1->Prev != NULL)
		P1 = P1->Prev;
	pt << P1 << P2;
}

void PutNodeBefore(TNode* newNode, TNode* node)
{
	if ( (newNode == NULL) || (node == NULL) )
		return;
	newNode->Next = node;
	newNode->Prev = node->Prev;
	node->Prev = newNode;
	if (newNode->Prev != NULL)
		newNode->Prev->Next = newNode;
}