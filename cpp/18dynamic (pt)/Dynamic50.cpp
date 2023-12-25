#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void MoveLast(TNode* node);
bool CanMoveRight(TNode* node);
void MoveRight(TNode* node);

void Solve()
{
    Task("Dynamic50");
	TNode* P1;
	pt >> P1;
	int counts = 0;
	for (TNode* i = P1; i != NULL; i = i->Next)
		counts++;
	TNode* node = P1;
	for (int i = 1; i <= counts; i++)
	{
		TNode* next = node->Next;
		if (node->Data % 2 != 0)
			MoveLast(node);
		node = next;
	}
	while (P1->Prev != NULL)
		P1 = P1->Prev;
	pt << P1;
}

void MoveLast(TNode* node)
{
	while (CanMoveRight(node))
		MoveRight(node);
}

bool CanMoveRight(TNode* node)
{
	if ( (node == NULL) || (node->Next == NULL) )
		return false;
	return true;
}

void MoveRight(TNode* node)
{
	if ( !CanMoveRight(node) )
		return;
	node->Next->Prev = node->Prev;
	if (node->Prev != NULL)
		node->Prev->Next = node->Next;
	node->Prev = node->Next;
	node->Next = node->Next->Next;
	node->Prev->Next = node;
	if (node->Next != NULL)
		node->Next->Prev = node;
}