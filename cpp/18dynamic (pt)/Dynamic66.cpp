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
	{
	}
};

TNode* DeleteCurrent(TList& L);
void AddAfterCurrent(TList& L, TNode* node);
void ToNext(TList& L);
bool IsLast(TList& L);
void SplitList(TList& L1, TList& L2);

void Solve()
{
    Task("Dynamic66");
	TNode* P1, *P2, *P3;
	pt >> P1 >> P2 >> P3;
	TList list1(P1, P2, P3);
	TList list2(NULL, NULL, NULL);
	SplitList(list1, list2);
	list1.Current = list1.First;
	list2.Current = list2.First;
	pt << list1.First << list1.Last << list1.Current;
	pt << list2.First << list2.Last << list2.Current;
}

void SplitList(TList& L1, TList& L2)
{
	TNode* node = NULL;
	while ( !IsLast(L1) )
	{
		node = DeleteCurrent(L1);
		AddAfterCurrent(L2, node);
	}
	node = DeleteCurrent(L1);
	AddAfterCurrent(L2, node);
}

TNode* DeleteCurrent(TList& L)
{
	if (L.Current == NULL)
		return NULL;
	TNode* next = L.Current->Next;
	if (L.Current->Next != NULL)
		L.Current->Next->Prev = L.Current->Prev;
	if (L.Current->Prev != NULL)
		L.Current->Prev->Next = L.Current->Next;
	
	if (L.Current == L.First)
	{
		L.First = L.First->Next;
		if (L.First != NULL)
			L.First->Prev = NULL;
	}
	if (L.Current == L.Last)
	{
		L.Last = L.Last->Prev;
		if (L.Last != NULL)
			L.Last->Next = NULL;
	}
	
	TNode* returnNode = L.Current;
	L.Current = next == NULL ? L.Last : next;
	
	return returnNode;
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
	L.Current = L.Last = node;
	if (L.First == NULL)
		L.First = L.Current;
}

void ToNext(TList& L)
{
	if ( (L.Current != NULL) && (L.Current->Next != NULL) )
		L.Current = L.Current->Next;
}

bool IsLast(TList& L)
{
	return (L.Current != NULL) && (L.Current == L.Last);
}