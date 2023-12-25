#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

bool MoveLeft(TNode* P);
bool MoveRight(TNode* P);

void Solve()
{
    Task("Dynamic48");
	TNode *Px, *Py;
	pt >> Px >> Py;
	int K = 0;
	for (TNode* ptr = Px; ptr != Py; ptr = ptr->Next)
		K++;
	for (int i = 1; i <= K; i++)
		MoveRight(Px);
	for (int i = 1; i < K; i++)
		MoveLeft(Py);
	TNode *P1 = Py;
	while (P1->Prev != NULL)
		P1 = P1->Prev;

	pt << P1;
}

bool MoveRight(TNode* P)
{
	if (P->Next == NULL)
		return false;

	P->Next->Prev = P->Prev;
	if (P->Prev != NULL)
		P->Prev->Next = P->Next;
	P->Prev = P->Next;
	P->Next = P->Prev->Next;
	P->Prev->Next = P;
	
	if (P->Next != NULL)
		P->Next->Prev = P;

	return true;
}

bool MoveLeft(TNode* P)
{
	if (P->Prev == NULL)
		return false;

	P->Prev->Next = P->Next;
	if (P->Next != NULL)
		P->Next->Prev = P->Prev;
	P->Next = P->Prev;
	P->Prev = P->Next->Prev;
	P->Next->Prev = P;
	
	if (P->Prev != NULL)
		P->Prev->Next = P;

	return true;
}