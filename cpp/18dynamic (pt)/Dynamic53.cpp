#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

TNode* GetFirst(TNode*);
TNode* GetLast(TNode*);

void DeleteThisNode(TNode*);
TNode* AddNodeAfter(TNode* node, TNode* newNode);

void Solve()
{
    Task("Dynamic53");
	TNode* Px, *Py;
	pt >> Px >> Py;

	TNode* P1 = GetFirst(Px);
	TNode *P2 = GetLast(Py);
	
	TNode* P3 = NULL, *P4 = NULL;
	TNode* notWith = Py->Next;
	for (TNode* i = Px; i != notWith; )
	{
		TNode* next = i->Next;
		if (P1 == i)
			P1 = i->Next;
		if (P2 == i)
			P2 = i->Prev;
		DeleteThisNode(i);
		P4 = AddNodeAfter(P4, i);
		i = next;
	}
	P4 = GetLast(P4);
	P3 = GetFirst(P4);

	pt << P1 << P3;
}

TNode* GetFirst(TNode* node)
{
	if (node == NULL)
		return NULL;
	TNode* first = node;
	while (first->Prev != NULL)
		first = first->Prev;
	return first;
}

TNode* GetLast(TNode* node)
{
	if (node == NULL)
		return NULL;
	TNode* last = node;
	while (last->Next != NULL)
		last = last->Next;
	return last;
}

void DeleteThisNode(TNode* node)
{
	if (node == NULL)
		return;
	if (node->Next != NULL)
		node->Next->Prev = node->Prev;
	if (node->Prev != NULL)
		node->Prev->Next = node->Next;
	node->Next = node->Prev = NULL;
}

TNode* AddNodeAfter(TNode* node, TNode* newNode)
{
	if (newNode == NULL)
		return NULL;
	newNode->Prev = node;
	newNode->Next = NULL;
	if (node != NULL)
	{
		newNode->Next = node->Next;
		node->Next = newNode;
		if (newNode->Next != NULL)
			newNode->Next->Prev = newNode;
	}
	return newNode;
}