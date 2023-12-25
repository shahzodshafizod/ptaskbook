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
	
	TNode* DeleteCurrent();
	void AddBeforeCurrent(TNode*);
	bool IsEmptyList() const;
	void InsertList(TList&);
	void ToFirst();
	void ToLast();
};

void Solve()
{
    Task("Dynamic68");
	TNode *P1, *P2, *P3, *P4, *P5, *P6;
	pt >> P1 >> P2 >> P3 >> P4 >> P5 >> P6;
	TList list1(P1, P2, P3);
	TList list2(P4, P5, P6);
	list2.InsertList(list1);
	pt << list2.First << list2.Last << list2.Current;
}

void TList::InsertList(TList& L)
{
	TNode* current = L.First;
	L.ToFirst();
	while ( !L.IsEmptyList() )
	{
		TNode* node = L.DeleteCurrent();
		this->AddBeforeCurrent(node);
		if (this->First == this->Current)
			this->First = node;
	}
	if (current != NULL)
		this->Current = current;
}

TNode* TList::DeleteCurrent()
{
	TNode* current = this->Current;
	if (this->Current == NULL)
		return NULL;
	this->Current = this->Current->Next;
	if (current->Prev != NULL)
		current->Prev->Next = current->Next;
	if (current->Next != NULL)
		current->Next->Prev = current->Prev;
	if (this->First == current)
		this->First = this->First->Next;
	if (this->Last == current)
		this->Last = this->Last->Prev;
	current->Next = current->Prev = NULL;
	return current;
}

void TList::AddBeforeCurrent(TNode* node)
{
	if (node == NULL)
		return;
	node->Next = this->Current;
	node->Prev = NULL;
	if (this->Current != NULL)
	{
		node->Prev = this->Current->Prev;
		this->Current->Prev = node;
		if (node->Prev != NULL)
			node->Prev->Next = node;
		if (this->First == this->Last)
			this->First = node;
	}
	else
		this->First = this->Last = this->Current = node;
}

bool TList::IsEmptyList() const
{
	return ( (this->Current == NULL) && (this->First == NULL) && (this->Last == NULL) );
}

void TList::ToFirst()
{
	this->Current = this->First;
}

void TList::ToLast()
{
	this->Current = this->Last;
}