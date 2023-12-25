#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

struct TQueue
{
	TNode* Head;
	TNode* Tail;
};

void Enqueue(TQueue& Q, int D);

void Solve()
{
    Task("Dynamic26");
	TNode* P1, *P2;
	pt >> P1 >> P2;
	int N, number;
	pt >> N;
	TQueue queue;
	queue.Head = P1;
	queue.Tail = P2;
	
	for (int i = 1; i <= N; i++)
	{
		pt >> number;
		Enqueue(queue, number);
	}

	pt << queue.Head << queue.Tail;
}

void Enqueue(TQueue& Q, int D)
{
	TNode* newTail = new TNode;
	newTail->Data = D;
	newTail->Next = NULL;
	if (Q.Tail == NULL)
		Q.Head = Q.Tail = newTail;
	else
	{
		Q.Tail->Next = newTail;
		Q.Tail = newTail;
	}
	newTail = NULL;
}