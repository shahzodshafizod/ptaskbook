#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

TNode* GetFirst(TNode* node);
TNode* GetLast(TNode* node);
void DeleteThisNode(TNode* node);
TNode* AddNodeAfter(TNode* node, TNode* newNode);

void Solve()
{
    Task("Dynamic54");
	TNode *Px, *Py;
	pt >> Px >> Py;
	TNode* P1 = GetFirst(Px);
	TNode* P2 = GetLast(Py);
	TNode* P4 = NULL;
	for (TNode* i = Px->Next; i != Py; )
	{
		TNode* next = i->Next;
		if (P1 == i)
			P1 = P1->Next;
		if (P2 == i)
			P2 = i->Prev;
		DeleteThisNode(i);
		P4 = AddNodeAfter(P4, i);
		i = next;
	}
	P4 = GetLast(P4);
	TNode* P3 = GetFirst(P4);

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
	while (last->Prev != NULL)
		last = last->Prev;
	return last;
}

void DeleteThisNode(TNode* node)
{
	if (node == NULL)
		return;
	if (node->Prev != NULL)
		node->Prev->Next = node->Next;
	if (node->Next != NULL)
		node->Next->Prev = node->Prev;
	node->Prev = node->Next = NULL;
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