#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

class TList
{
public:
	TNode* First;
	TNode* Last;
	TNode* Current;

	TList(TNode* P1, TNode* P2, TNode* P3)
		: First(P1), Last(P2), Current(P3)
	{}
};

void MoveCurrent(TList& L1, TList& L2);
TNode* DeleteCurrent(TList& L);
void AddAfterCurrent(TList& L, TNode* node);

void Solve()
{
    Task("Dynamic69");
	TNode *P1, *P2, *P3, *P4, *P5, *P6;
	pt >> P1 >> P2 >> P3 >> P4 >> P5 >> P6;
	TList list1(P1, P2, P3);
	TList list2(P4, P5, P6);
	MoveCurrent(list1, list2);
	pt << list1.First << list1.Last << list1.Current;
	pt << list2.First << list2.Last << list2.Current;
}

void MoveCurrent(TList& L1, TList& L2)
{
	TNode* node = DeleteCurrent(L1);
	AddAfterCurrent(L2, node);
}

TNode* DeleteCurrent(TList& L)
{
	TNode* current = L.Current;
	if (current == NULL)
		return NULL;
	L.Current = L.Current->Next;
	if (current->Next != NULL)
		current->Next->Prev = current->Prev;
	if (current->Prev != NULL)
		current->Prev->Next = current->Next;
	if (L.First == current)
		L.First = L.First->Next;
	if (L.Last == current)
		L.Last = L.Current = L.Last->Prev;
	current->Next = current->Prev = NULL;
	return current;
}

void AddAfterCurrent(TList& L, TNode* node)
{
	if (node == NULL)
		return;
	node->Prev = L.Current;
	node->Next = NULL;
	if (L.Current != NULL)
	{
		node->Next = L.Current->Next;
		L.Current->Next = node;
		if (node->Next != NULL)
			node->Next->Prev = node;
	}
	else
		L.First = L.Last = node;
	if (L.Current == L.Last)
		L.Last = node;
	L.Current = node;
}