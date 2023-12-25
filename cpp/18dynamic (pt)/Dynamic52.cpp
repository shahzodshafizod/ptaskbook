#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void PutNodeAfter(TNode* node, TNode* newNode);

void Solve()
{
    Task("Dynamic52");
	TNode *P1, *P2, *P0;
	pt >> P1 >> P2 >> P0;
	while (P2 != NULL)
	{
		TNode* prev = P2->Prev;
		PutNodeAfter(P0, P2);
		P2 = prev;
	}
	P1 = P0;
	while (P1->Prev != NULL)
		P1 = P1->Prev;
	P2 = P0;
	while (P2->Next != NULL)
		P2 = P2->Next;
	pt << P1 << P2;
}

void PutNodeAfter(TNode* node, TNode* newNode)
{
	if (newNode == NULL)
		return;
	newNode->Prev = node;
	newNode->Next = node->Next;
	node->Next = newNode;
	if (newNode->Next != NULL)
		newNode->Next->Prev = newNode;
}