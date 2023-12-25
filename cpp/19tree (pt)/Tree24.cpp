#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

void FindMaxUneven(TNode* P, int& maximal, bool& initialized);
TNode* GetLastMaxNode(TNode* P, int maximal);

void Solve()
{
    Task("Tree24");
	TNode* P1;
	pt >> P1;
	int maximal;
	bool initialized = false;
	FindMaxUneven(P1, maximal, initialized);
	TNode* P2 = initialized == false ? NULL : GetLastMaxNode(P1, maximal);
	pt << P2;
}

void FindMaxUneven(TNode* P, int& maximal, bool& initialized)
{
	if (P == NULL)
		return;
	FindMaxUneven(P->Left, maximal, initialized);
	int data = P->Data;
	if (data % 2 != 0)
	{
		if (initialized == false)
		{
			maximal = data;
			initialized = true;
		}
		else if (data > maximal)
			maximal = data;
	}
	FindMaxUneven(P->Right, maximal, initialized);
}

TNode* GetLastMaxNode(TNode* P, int maximal)
{
	if (P == NULL)
		return NULL;
	TNode* node = NULL;
	
	TNode* left = GetLastMaxNode(P->Left, maximal);
	if (left != NULL) node = left;
	
	if (P->Data == maximal)
		node = P;
	
	TNode* right = GetLastMaxNode(P->Right, maximal);
	if (right != NULL) node = right;

	return node;
}