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

void ToLast(TList& L);
void ToPrev(TList& L);
int GetData(const TList& L);
bool IsFirst(const TList& L);

void Solve()
{
    Task("Dynamic64");
	TNode *P1, *P2, *P3;
	pt >> P1 >> P2 >> P3;
	TList list = {P1, P2, P3};
	/*
	list.First = P1;
	list.Last = P2;
	list.Current = P3;
	*/
	int data, counts = 1;
	ToLast(list);
	while ( !IsFirst(list) )
	{
		data = GetData(list);
		if (data % 2 == 0)
			pt << data;
		counts++;
		ToPrev(list);
		if (IsFirst(list)) break;
		data = GetData(list);
		if (data % 2 == 0)
			pt << data;
		counts++;
		ToPrev(list);
	}
	data = GetData(list);
	if (data % 2 == 0)
		pt << data;
	pt << counts;
}

void ToLast(TList& L)
{
	L.Current = L.Last;
}

void ToPrev(TList& L)
{
	if ( (L.Current != NULL) && (L.Current->Prev != NULL) )
		L.Current = L.Current->Prev;
}

int GetData(const TList& L)
{
	return L.Current->Data;
}

bool IsFirst(const TList& L)
{
	return L.Current == L.First;
}