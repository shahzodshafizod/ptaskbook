#ifndef ____STACK____H____
#define ____STACK____H____
#include "TNode.h"
#include <stdbool.h>

struct TStack {
	struct TNode* Top;
};

void TStack_constr(struct TStack*);
void TStack_destr(struct TStack*);
void TStack_handmake(struct TStack*);
void TStack_automake(struct TStack*, int);
void TStack_free(struct TStack*);
void TStack_Push(struct TStack*, int);
int TStack_Pop(struct TStack*);
bool TStack_IsEmpty(const struct TStack*);
int TStack_Peek(const struct TStack*);
struct TNode* TStack_getTop(const struct TStack*);
void TStack_print(const struct TStack*);

int TStack_getRandomData(int, int);

#endif