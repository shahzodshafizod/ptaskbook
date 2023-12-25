#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

struct TList
{
	TNode* First;
	TNode* Last;
};

TNode* DeleteThisNode(TNode* node);
void AddAfter(TNode** node, TNode* newNode);

void Solve()
{
    Task("Dynamic71");
	TNode* P1, *P2;
	pt >> P1 >> P2;
	TNode* P3 = new TNode;
	P3->Data = 0;
	TNode *P4 = P3;
	if (P1 != P2)
	{
		TNode* lastInFirstList = P2->Prev;
		while (P2 != P1)
		{
			TNode* node = P2;
			P2 = P2->Next;
			node = DeleteThisNode(node);
			AddAfter(&P4, node);
		}
		P2 = lastInFirstList;
		if (P1 == P2)
			P1->Next = P1->Prev = P1;
		else
		{
			P1->Prev = P2;
			P2->Next = P1;
		}
	}
	if (P4 == P3)
		P3->Next = P3->Prev = P3;
	else
	{
		P3->Prev = P4;
		P4->Next = P3;
	}
	pt << P3;
}

TNode* DeleteThisNode(TNode* node)
{
	if (node == NULL)
		return NULL;
	if (node->Next != NULL)
		node->Next->Prev = node->Prev;
	if (node->Prev != NULL)
		node->Prev->Next = node->Next;
	node->Next = node->Prev = NULL;
	return node;
}

void AddAfter(TNode** node, TNode* newNode)
{
	if (newNode == NULL)
		return;
	newNode->Prev = *node;
	newNode->Next = NULL;
	if (*node != NULL)
	{
		newNode->Next = (*node)->Next;
		(*node)->Next = newNode;
		if ((*node)->Next != NULL)
			(*node)->Next->Prev = *node;
	}
	*node = newNode;
}