#include "TNode.h"

TNode::TNode() {
	this->Data = 0;
	this->Parent = this->Left = this->Right = 0;
}