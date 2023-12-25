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

void ToFirst(TList&);
void ToNext(TList&);
void SetData(TList&, int);
bool IsLast(const TList&);

void Solve()
{
    Task("Dynamic63");
	TNode* P1, *P2, *P3;
	pt >> P1 >> P2 >> P3;
	TList list = {P1, P2, P3};
	/*
	list.First = P1;
	list.Last = P2;
	list.Current = P3;
	*/
	int counts = 1;
	for (ToFirst(list); !IsLast(list); )
	{
		list.Current->Data = 0;
		
		ToNext(list);
		counts++;
		if (IsLast(list)) break;
		
		ToNext(list);
		counts++;
	}
	if (counts % 2 != 0)
		list.Last->Data = 0;
	
	pt << counts;
	pt << list.First << list.Last << list.Current;
}

void ToFirst(TList& L)
{
	L.Current = L.First;
}

void ToNext(TList& L)
{
	if ( (L.Current != NULL) && (L.Current->Next != NULL) )
		L.Current = L.Current->Next;
}

void SetData(TList& L, int D)
{
	if (L.Current != NULL)
		L.Current->Data = D;
}

bool IsLast(const TList& L)
{
	return L.Current == L.Last;
}