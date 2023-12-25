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

	TList(TNode *p1, TNode *p2, TNode *p3)
		: First(p1), Last(p2), Current(p3)
	{}
};

void AddList(TList& L1, TList& L2);
void ToFirst(TList& L);
bool IsEmpty(TList& L);
TNode* DeleteCurrent(TList& L);
void AddAtLast(TList& L, TNode* node);

void Solve()
{
    Task("Dynamic67");
	TNode *P1, *P2, *P3, *P4, *P5, *P6;
	pt >> P1 >> P2 >> P3 >> P4 >> P5 >> P6;
	TList list1(P1, P2, P3);
	TList list2(P4, P5, P6);
	AddList(list1, list2);
	pt << list2.First << list2.Last << list2.Current;
}

void AddList(TList& L1, TList& L2)
{
	TNode* current = L1.First;
	for (ToFirst(L1); !IsEmpty(L1); )
	{
		TNode* node = DeleteCurrent(L1);
		AddAtLast(L2, node);
	}
	if (current != NULL)
		L2.Current = current;
}

void ToFirst(TList& L)
{
	L.Current = L.First;
}

bool IsEmpty(TList& L)
{
	return L.Current == NULL;
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
		L.Last = L.Last->Prev;
	current->Next = current->Prev = NULL;
	return current;
}

void AddAtLast(TList& L, TNode* node)
{
	if (node == NULL)
		return;
	node->Prev = L.Last;
	node->Next = NULL;
	if (L.Last != NULL)
		L.Last->Next = node;
	L.Last = node;
}