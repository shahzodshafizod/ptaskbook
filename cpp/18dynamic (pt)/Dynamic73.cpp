#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

TNode* DeleteThisNode(TNode *node);
void AddAfter(TNode* node, TNode* newNode);

void Solve()
{
    Task("Dynamic73");
	TNode* P1, *P2;
	pt >> P1 >> P2;
	TNode* P1Last = P1->Prev;
	TNode* P2Last = P2->Prev;
	bool IsSecondsBarer = P2 == P2Last;
	while (P1 != P1Last)
	{
		TNode* node = P1Last;
		P1Last = P1Last->Prev;
		node = DeleteThisNode(node);
		AddAfter(P2, node);
	}
	delete P1;
	if (IsSecondsBarer)
		P2Last = P2->Prev;
	P2 = P2->Next;
	pt << P2 << P2Last;
}

TNode* DeleteThisNode(TNode *node)
{
	if (node != NULL)
	{
		if (node->Next != NULL)
			node->Next->Prev = node->Prev;
		if (node->Prev != NULL)
			node->Prev->Next = node->Next;
		node->Next = node->Prev = NULL;
	}
	return node;
}

void AddAfter(TNode* node, TNode* newNode)
{
	if (newNode == NULL)
		return;
	newNode->Next = NULL;
	newNode->Prev = node;
	if (node != NULL)
	{
		newNode->Next = node->Next;
		node->Next = newNode;
		if (newNode->Next != NULL)
			newNode->Next->Prev = newNode;
	}
	else
		node = newNode;
}