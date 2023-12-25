#include <windows.h>
#pragma hdrstop
#include "pt4.h"
using namespace std;

struct TQueue
{
	TNode* Head;
	TNode* Tail;
};

bool QueueIsEmpty(const TQueue& queue);
int Dequeue(TQueue& queue);

void Solve()
{
    Task("Dynamic28");
	TNode* P1, *P2;
	pt >> P1 >> P2;
	TQueue queue;
	queue.Head = P1;
	queue.Tail = P2;
	for (int i = 0; (i < 5) && !QueueIsEmpty(queue); i++)
		pt << Dequeue(queue);
	pt << QueueIsEmpty(queue);
	pt << queue.Head << queue.Tail;
}

bool QueueIsEmpty(const TQueue& queue)
{
	if (queue.Head == NULL)
		return true;
	return false;
}

int Dequeue(TQueue& queue)
{
	TNode* temp = queue.Head;
	queue.Head = queue.Head->Next;
	if (queue.Head == NULL) queue.Tail = NULL;
	int D = temp->Data;
	delete temp;
	return D;
}