#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

struct TList
{
	TNode* First;
	TNode* Last;
	TNode* Current;
};

void InsertAfter(TList& L, int D);

void Solve()
{
    Task("Dynamic62");
	TNode *P1, *P2, *P3;
	pt >> P1 >> P2 >> P3;
	TList list = {P1, P2, P3};
	/*
	list.First = P1;
	list.Last = P2;
	list.Current = P3;
	*/
	int number;
	for (int i = 1; i <= 5; i++)
	{
		pt >> number;
		InsertAfter(list, number);
	}
	while (list.Last->Next != NULL)
		list.Last = list.Last->Next;
	pt << list.First << list.Last << list.Current;
}

void InsertAfter(TList& L, int D)
{
	TNode* node = new TNode;
	node->Data = D;
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
	L.Current = node;
}