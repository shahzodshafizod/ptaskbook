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

void InsertBefore(TList& L, int D);

void Solve()
{
    Task("Dynamic61");
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
		InsertBefore(list, number);
	}
	while (list.First->Prev != NULL)
		list.First = list.First->Prev;
	pt << list.First << list.Last << list.Current;
}

void InsertBefore(TList& L, int D)
{
	TNode* node = new TNode;
	node->Data = D;
	node->Next = L.Current;
	node->Prev = NULL;
	if (L.Current != NULL)
	{
		node->Prev = L.Current->Prev;
		L.Current->Prev = node;
		if (node->Prev != NULL)
			node->Prev->Next = node;
	}
	else
		L.First = L.Last = node;
	L.Current = node;
}