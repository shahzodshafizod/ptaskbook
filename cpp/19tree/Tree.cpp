#include "Tree.h"
#include <iostream>
using namespace std;

Tree::Tree() {
	this->root = NULL;
	this->field = NULL;
	this->nodeCount = 0;
	this->level = -1;
}

Tree::Tree(Tree& rh) {
	this->root = NULL;
	this->field = NULL;
	this->nodeCount = 0;
	this->level = -1;
	this->copy(rh);
}

Tree::~Tree() {
	if (this->current != this->root) {
		this->current = this->root;
	}
	this->free();
	delete this->field;
}

void Tree::make() {
	if (this->root == NULL) {
		this->root = new TNode();
		this->nodeCount++;
		cout << "Root's data: ";
		cin >> this->root->Data;
		this->current = this->root;
	}
	int answer;
	cout << "Where to go? (0-exit; 1-left; 2-right; 3-parent):\t";
	cin >> answer;
	if (answer == 0) { this->current = this->root; return; }
	if (answer > 3 || answer < 0) { this->make(); return; }
	if (answer == 3) { this->current = this->current->Parent; this->make(); return; }
	TNode* newLeaf = new TNode();
	this->nodeCount++;
	cout << "data:\t";
	cin >> newLeaf->Data;
	newLeaf->Parent = this->current;
	switch (answer) {
		case 1: this->current->Left = newLeaf; break;
		case 2: this->current->Right = newLeaf; break;
	}
	this->current = newLeaf;
	this->make();
}

void Tree::make(int arr[], int index, int N, ChildNode child) {
	if (index >= N) {
		this->current = this->root;
		return;
	}
	if (this->root == NULL) {
		this->root = new TNode();
		this->nodeCount++;
		this->root->Data = arr[index++];
		this->current = this->root;
		if (index >= N) { return; }
	}
	TNode* newLeaf = new TNode();
	this->nodeCount++;
	newLeaf->Data = arr[index++];
	newLeaf->Parent = this->current;
	switch (child) {
		case TREE25:
			this->current->Right = newLeaf;
			break;
		case TREE26LEFT:
			this->current->Left = newLeaf;
			child = TREE26RIGHT;
			break;
		case TREE26RIGHT:
			this->current->Right = newLeaf;
			child = TREE26LEFT;
			break;
		case TREE27:
			if (this->current->Data % 2 != 0) {
				this->current->Left = newLeaf;
			} else {
				this->current->Right = newLeaf;
			}
			break;
		case TREE28:
			this->current->Left = newLeaf;
			if (index < N) {
				newLeaf = new TNode();
				this->nodeCount++;
				newLeaf->Data = arr[index++];
				newLeaf->Parent = this->current;
				this->current->Right = newLeaf;
			}
			break;
		case TREE29:
			TNode* tempNode = NULL;
			if (index < N) {
				tempNode = new TNode();
				this->nodeCount++;
				tempNode->Data = arr[index++];
				tempNode->Parent = this->current;
			}
			if (this->current->Data % 2 != 0) {
				this->current->Left = newLeaf;
				this->current->Right = tempNode;
			} else {
				this->current->Right = newLeaf;
				this->current->Left = tempNode;
			}
			newLeaf = tempNode;
			break;
	}
	this->current = newLeaf;
	this->make(arr, index, N, child);
}

void Tree::make(int arr[], int& index, int N, int level, int L, bool toLeft) {
	if (index >= N || level > L) {
		this->current = this->root;
		return;
	}
	TNode* newLeaf = new TNode();
	this->nodeCount++;
	newLeaf->Data = arr[index++];
	newLeaf->Parent = this->current;
	if (this->root == NULL) {
		this->root = this->current = newLeaf;
	} else {
		switch (toLeft) {
			case true: this->current->Left = newLeaf; break;
			case false: this->current->Right = newLeaf; break;
		}
	}
	this->current = newLeaf;
	this->make(arr, index, N, level+1, L, true);
	this->current = newLeaf;
	this->make(arr, index, N, level+1, L, false);
}

void Tree::make(int arr[], int& index, int N, int count, bool toLeft) {
	if (index >= N || count <= 0) {
		this->current = this->root;
		return;
	}
	TNode* newLeaf = new TNode();
	this->nodeCount++;
	newLeaf->Data = arr[index++];
	newLeaf->Parent = this->current;
	if (this->root == NULL) {
		this->current = this->root = newLeaf;
	} else {
		switch (toLeft) {
			case true: this->current->Left = newLeaf; break;
			case false: this->current->Right = newLeaf; break;
		}
	}
	this->current = newLeaf;
	this->make(arr, index, N, count/2, true);
	this->current = newLeaf;
	this->make(arr, index, N, count - 1 - count/2, false);
}

void Tree::make(int level, int count, bool toLeft) {
	if (count <= 0) {
		this->current = this->root;
		return;
	}
	TNode* newLeaf = new TNode();
	this->nodeCount++;
	newLeaf->Data = level;
	newLeaf->Parent = this->current;
	if (this->root == NULL) {
		this->current = this->root = newLeaf;
	} else {
		switch (toLeft) {
			case true: this->current->Left = newLeaf; break;
			case false: this->current->Right = newLeaf; break;
		}
	}
	this->current = newLeaf;
	this->make(level+1, count/2, true);
	this->current = newLeaf;
	this->make(level+1, count - 1 - count/2, false);
}

void Tree::make(int K) {
	if (K < 1) {
		this->current = this->root;
		return;
	}
	TNode* newLeaf = new TNode();
	this->nodeCount++;
	newLeaf->Data = K;
	newLeaf->Parent = this->current;
	if (this->root == NULL) {
		this->current = this->root = newLeaf;
	} else if (2*K > this->current->Data) {
		this->current->Right = newLeaf;
	} else {
		this->current->Left = newLeaf;
	}
	
	if (K == 1) {
		this->current = this->root;
		return;
	}
	this->current = newLeaf;
	this->make(K/2);
	if (K % 2 != 0) {
		this->current = newLeaf;
		this->make(K - K/2);
	}
}

void Tree::copy(Tree& rh, bool toRight) {
	if (rh.current == NULL) {
		rh.current = rh.root;
		this->current = this->root;
		return;
	}
	TNode* newLeaf = new TNode();
	this->nodeCount++;
	newLeaf->Data = rh.current->Data;
	newLeaf->Parent = this->current;
	if (this->root == NULL) {
		this->current = this->root = newLeaf;
	} else {
		switch (toRight) {
			case false: this->current->Left = newLeaf; break;
			case true: this->current->Right = newLeaf; break;
		}
	}
	TNode* node = rh.current;
	
	this->current = newLeaf;
	rh.current = node->Left;
	this->copy(rh);
	
	this->current = newLeaf;
	rh.current = node->Right;
	this->copy(rh, true);
}

void Tree::copy(TNode* node, bool toRight) {
	if (node == NULL) {
		this->current = this->root;
		return;
	}
	TNode* newLeaf = new TNode();
	this->nodeCount++;
	newLeaf->Data = node->Data;
	newLeaf->Parent = this->current;
	if (this->root == NULL) {
		this->current = this->root = newLeaf;
	} else {
		switch (toRight) {
			case true: this->current->Right = newLeaf; break;
			case false: this->current->Left = newLeaf; break;
		}
	}
	this->current = newLeaf;
	this->copy(node->Left);
	this->current = newLeaf;
	this->copy(node->Right, true);
}

void Tree::free() {
	if (this->current == NULL) {
		this->current = this->root;
		return;
	}
	TNode* node = this->current;
	
	this->current = node->Left;
	this->free();
	this->current = node->Right;
	this->free();
	
	this->current = node;
	delete this->current;
	this->nodeCount--;
	this->current = NULL;
}

void Tree::display(TNode* node) {
	if (this->root == NULL) {
		return;
	}
	if (this->field == NULL) {
		this->field = new Field(this->root, this->getLevel()+1, this->getNodeCount()+1);
	}
	field->display(node);
}

void Tree::infix() {
	if (this->current == NULL) {
		this->current = this->root;
		return;
	}
	TNode* node = this->current;
	this->current = node->Left;
	this->infix();
	cout << node->Data << '\t';
	this->current = node->Right;
	this->infix();
}

void Tree::prefix() {
	if (this->current == NULL) {
		this->current = this->root;
		return;
	}
	TNode* node = this->current;
	cout << node->Data << '\t';
	this->current = node->Left;
	this->prefix();
	this->current = node->Right;
	this->prefix();
}

void Tree::postfix() {
	if (this->current == NULL) {
		this->current = this->root;
		return;
	}
	TNode* node = this->current;
	this->current = node->Left;
	this->postfix();
	this->current = node->Right;
	this->postfix();
	cout << node->Data << '\t';
}

void Tree::infixToN(int &index, int N) {
	if (this->current == NULL) {
		this->current = this->root;
		return;
	}
	TNode* node = this->current;
	this->current = node->Left;
	this->infixToN(index, N);
	++index;
	if (index <= N) {
		cout << node->Data << '\t';
	} else {
		this->current = this->root;
		return;
	}
	this->current = node->Right;
	this->infixToN(index, N);
}

void Tree::postfixFromN(int& index, int N) {
	if (this->current == NULL) {
		this->current = this->root;
		return;
	}
	TNode* node = this->current;
	this->current = node->Left;
	this->postfixFromN(index, N);
	this->current = node->Right;
	this->postfixFromN(index, N);
	++index;
	if (index >= N) {
		cout << node->Data << '\t';
	}
}

void Tree::prefixBetween(int& index, int N1, int N2) {
	if (this->current == NULL) {
		this->current = this->root;
		return;
	}
	TNode* node = this->current;
	++index;
	if (N1 <= index && index < N2) {
		cout << node->Data << '\t';
	} else if (index > N2) {
		this->current = this->root;
		return;
	}
	this->current = node->Left;
	this->prefixBetween(index, N1, N2);
	this->current = node->Right;
	this->prefixBetween(index, N1, N2);
}

int Tree::getNodeCount() const {
	return this->nodeCount;
}

int Tree::getLeftCount(bool isLeft) {
	if (this->current == NULL) {
		this->current = this->root;
		return 0;
	}
	int count = isLeft ? 1 : 0;
	TNode* node = this->current;
	this->current = node->Left;
	count += this->getLeftCount(true);
	this->current = node->Right;
	count += this->getLeftCount();
	return count;
}

int Tree::getLeafCount() {
	if (this->current == NULL) {
		this->current = this->root;
		return 0;
	}
	int count = this->current->Left == NULL && this->current->Right == NULL ? 1 : 0;
	TNode* node = this->current;
	this->current = node->Left;
	count += this->getLeafCount();
	this->current = node->Right;
	count += this->getLeafCount();
	return count;
}

int Tree::getRightLeafCount(bool isRight) {
	if (this->current == NULL) {
		this->current = this->root;
		return 0;
	}
	int count = 0;
	if (isRight && this->current->Left == NULL && this->current->Right == NULL) {
		count = 1;
	}
	TNode* node = this->current;
	this->current = node->Left;
	count += this->getRightLeafCount();
	this->current = node->Right;
	count += this->getRightLeafCount(true);
	return count;
}

int Tree::getNodeCountK(int K) {
	if (this->current == NULL) {
		this->current = this->root;
		return 0;
	}
	TNode* node = this->current;
	int count = node->Data == K ? 1 : 0;
	this->current = node->Left;
	count += this->getNodeCountK(K);
	this->current = node->Right;
	count += this->getNodeCountK(K);
	return count;
}

int Tree::getLeafCountK(int K) {
	if (this->current == NULL) {
		this->current = this->root;
		return 0;
	}
	TNode* node = this->current;
	int count = 0;
	if (node->Data == K && node->Left == NULL && node->Right == NULL) {
		count = 1;
	}
	this->current = node->Left;
	count += this->getLeafCountK(K);
	this->current = node->Right;
	count += this->getLeafCountK(K);
	return count;
}

void Tree::setLevel(int currentLevel) {
	if (this->current == NULL) {
		this->current = this->root;
		return;
	}
	if (currentLevel > this->level) {
		this->level = currentLevel;
	}
	TNode* node = this->current;
	this->current = node->Left;
	this->setLevel(currentLevel+1);
	this->current = node->Right;
	this->setLevel(currentLevel+1);
}

int Tree::getLevel() {
	if (this->level < 0) {
		this->setLevel();
	}
	return this->level;
}

int Tree::getLevelNodeCount(int L, int level)  {
	if (this->current == NULL) {
		this->current = this->root;
		return 0;
	}
	TNode* node = this->current;
	this->current = node->Left;
	int count = this->getLevelNodeCount(L, level+1);
	if (level == L) {
		cout << node->Data << '\t';
		++count;
	} else if (level > L) {
		this->current = this->root;
		return count;
	}
	this->current = node->Right;
	count += this->getLevelNodeCount(L, level+1);
	return count;
}

int Tree::getDataSum() {
	if (this->current == NULL) {
		this->current = this->root;
		return 0;
	}
	int sum = this->current->Data;
	TNode* node = this->current;
	this->current = node->Left;
	sum += this->getDataSum();
	this->current = node->Right;
	sum += this->getDataSum();
	return sum;
}

int Tree::getLeafDataSum() {
	if (this->current == NULL) {
		this->current = this->root;
		return 0;
	}
	int sum = 0;
	if (this->current->Left == NULL && this->current->Right == NULL) {
		sum = this->current->Data;
	}
	TNode* node = this->current;
	this->current = node->Left;
	sum += this->getLeafDataSum();
	this->current = node->Right;
	sum += this->getLeafDataSum();
	return sum;
}

void Tree::levelNodeCountToArr(int arr[], int index) {
	if (this->current == NULL) {
		this->current = this->root;
		return;
	}
	arr[index]++;
	TNode* node = this->current;
	this->current = node->Left;
	this->levelNodeCountToArr(arr, index+1);
	this->current = node->Right;
	this->levelNodeCountToArr(arr, index+1);
}

void Tree::levelNodeSumToArr(int arr[], int index) {
	if (this->current == NULL) {
		this->current = this->root;
		return;
	}
	arr[index] += this->current->Data;
	TNode* node = this->current;
	this->current = node->Left;
	this->levelNodeSumToArr(arr, index+1);
	this->current = node->Right;
	this->levelNodeSumToArr(arr, index+1);
}

int Tree::getMaxData() {
	int maximal;
	if (this->current->Left == NULL && this->current->Right == NULL) {
		maximal = this->current->Data;
		this->current = this->root;
		return maximal;
	}
	TNode* node = this->current;
	int data;
	maximal = node->Data;
	if (node->Left != NULL) {
		this->current = node->Left;
		data = this->getMaxData();
		if (data > maximal) {
			maximal = data;
		}
	}
	if (node->Right != NULL) {
		this->current = node->Right;
		data = this->getMaxData();
		if (data > maximal) {
			maximal = data;
		}
	}
	return maximal;
}

int Tree::getMinData() {
	int minimal;
	if (this->current->Left == NULL && this->current->Right == NULL) {
		minimal = this->current->Data;
		this->current = this->root;
		return minimal;
	}
	TNode* node = this->current;
	int data;
	minimal = node->Data;
	if (node->Left != NULL) {
		this->current = node->Left;
		data = this->getMinData();
		if (data < minimal) {
			minimal = data;
		}
	}
	if (node->Right != NULL) {
		this->current = node->Right;
		data = this->getMinData();
		if (data < minimal) {
			minimal = data;
		}
	}
	return minimal;
}

int Tree::getMinLeafData() {
	int minimal;
	if (this->current->Left == NULL && this->current->Right == NULL) {
		minimal = this->current->Data;
		this->current = this->root;
		return minimal;
	}
	TNode* node = this->current;
	int data;
	bool inited = false;
	if (node->Left != NULL) {
		this->current = node->Left;
		minimal = this->getMinLeafData();
		inited = true;
	}
	if (node->Right != NULL) {
		this->current = node->Right;
		data = this->getMinLeafData();
		if (!inited) {
			minimal = data;
		} else if (data < minimal) {
			minimal = data;
		}
	}
	return minimal;
}

int Tree::getMaxInternalData() {
	if (this->current->Left == NULL && this->current->Right == NULL) {
		this->current = this->root;
		return 0;
	}
	TNode* node = this->current;
	int data, maximal = node->Data;
	if (node->Left != NULL && (node->Left->Left != NULL || node->Left->Right != NULL)) {
		this->current = node->Left;
		data = this->getMaxInternalData();
		if (data > maximal) {
			maximal = data;
		}
	}
	if (node->Right != NULL && (node->Right->Left != NULL || node->Right->Right != NULL)) {
		this->current = node->Right;
		data = this->getMaxInternalData();
		if (data > maximal) {
			maximal = data;
		}
	}
	return maximal;
}

TNode* Tree::getFirstNodePrefix(int data) {
	if (this->current == NULL) {
		this->current = this->root;
		return NULL;
	}
	TNode* node = this->current;
	if (node->Data == data) {
		return node;
	}
	this->current = node->Left;
	TNode* tempNode = this->getFirstNodePrefix(data);
	if (tempNode == NULL) {
		this->current = node->Right;
		tempNode = this->getFirstNodePrefix(data);
	}
	return tempNode;
}

TNode* Tree::getLastNodeInfix(int data) {
	if (this->current == NULL) {
		this->current = this->root;
		return NULL;
	}
	TNode* node = this->current;
	this->current = node->Left;
	TNode* resNode = this->getLastNodeInfix(data);
	if (node->Data == data) {
		resNode = node;
	}
	this->current = node->Right;
	TNode* tempNode = this->getLastNodeInfix(data);
	if (tempNode != NULL) {
		resNode = tempNode;
	}
	return resNode;
}

bool Tree::hasOddData() {
	if (this->current == NULL) {
		this->current = this->root;
		return false;
	}
	TNode* node = this->current;
	if (node->Data % 2 != 0) {
		return true;
	}
	this->current = node->Left;
	bool hasIt = this->hasOddData();
	if (!hasIt) {
		this->current = node->Right;
		hasIt = this->hasOddData();
	}
	return hasIt;
}

int Tree::getMaxOddData() {
	int maximal;
	if (this->current->Left == NULL && this->current->Right == NULL) {
		maximal = this->current->Data;
		this->current = this->root;
		return maximal;
	}
	TNode* node = this->current;
	int data;
	bool inited = false;
	if (node->Data % 2 != 0) {
		maximal = node->Data;
		inited = true;
	}
	if (node->Left != NULL) {
		this->current = node->Left;
		data = this->getMaxOddData();
		if (!inited) {
			maximal = data;
			inited = true;
		} else if (data > maximal) {
			maximal = data;
		}
	}
	if (node->Right != NULL) {
		this->current = node->Right;
		data = this->getMaxOddData();
		if (!inited) {
			maximal = data;
			inited = true;
		} else if (data > maximal) {
			maximal = data;
		}
	}
	return maximal;
}

void Tree::changeDatas(ChildNode child) {
	if (this->current == NULL) {
		this->current = this->root;
		return;
	}
	TNode* node = this->current;
	
	switch (child) {
		case TREE35:
			node->Data *= 2;
			break;
		case TREE36:
			node->Data /= node->Data % 2 == 0 ? 2 : 1;
			break;
		case TREE37:
			node->Data += node->Left == NULL && node->Right == NULL ? 1 : -1;
			break;
		case TREE38:
			if (node->Left != NULL && node->Right != NULL) {
				int data = node->Left->Data;
				node->Left->Data = node->Right->Data;
				node->Right->Data = data;
			}
			break;
	}
	this->current = node->Left;
	this->changeDatas(child);
	this->current = node->Right;
	this->changeDatas(child);
}

void Tree::swapChilds() {
	if (this->current == NULL) {
		this->current = this->root;
		delete this->field;
		this->field = NULL;
		return;
	}
	TNode* node = this->current->Left;
	this->current->Left = this->current->Right;
	this->current->Right = node;
	
	node = this->current;
	this->current = node->Left;
	this->swapChilds();
	this->current = node->Right;
	this->swapChilds();
}

void Tree::deleteAllChilds() {
	if (this->current == NULL) {
		this->current = this->root;
		delete this->field;
		this->field = NULL;
		this->level = -1;
		return;
	}
	TNode* node = this->current;
	this->current = node->Left;
	this->deleteAllChilds();
	this->current = node->Right;
	this->deleteAllChilds();
	node->Left = node->Right = NULL;
	if (node != this->root) {
		delete node;
		this->nodeCount--;
	}
}

void Tree::deleteAllLeaves() {
	if (this->current == NULL) {
		this->current = this->root;
		delete this->field;
		this->field = NULL;
		this->level = -1;
		return;
	}
	TNode* node = this->current;
	if (node->Left != NULL && node->Left->Left == NULL && node->Left->Right == NULL) {
		delete node->Left;
		this->nodeCount--;
		node->Left = NULL;
	}
	if (node->Right != NULL && node->Right->Left == NULL && node->Right->Right == NULL) {
		delete node->Right;
		this->nodeCount--;
		node->Right = NULL;
	}
	this->current = node->Left;
	this->deleteAllLeaves();
	this->current = node->Right;
	this->deleteAllLeaves();
}

void Tree::deleteSubTree(ChildNode child) {
	if (this->current == NULL) {
		this->current = this->root;
		delete this->field;
		this->field = NULL;
		this->level = -1;
		return;
	}
	TNode* node = this->current;
	switch (child) {
		case TREE42:
			if (node->Left != NULL && node->Left->Data < this->root->Data) {
				this->current = node->Left;
				this->free();
				node->Left = NULL;
			}
			if (node->Right != NULL && node->Right->Data < this->root->Data) {
				this->current = node->Right;
				this->free();
				node->Right = NULL;
			}
			break;
		case TREE43:
			if (node->Left != NULL && node->Right != NULL) {
				if (node->Data % 2 == 0) {
					this->current = node->Right;
					this->free();
					node->Right = NULL;
				} else {
					this->current = node->Left;
					this->free();
					node->Left = NULL;
				}
			}
			break;
	}
	this->current = node->Left;
	this->deleteSubTree(child);
	this->current = node->Right;
	this->deleteSubTree(child);
}

void Tree::addLeaves(ChildNode child) {
	if (this->current == NULL) {
		this->current = this->root;
		delete this->field;
		this->field = NULL;
		this->level = -1;
		return;
	}
	TNode* node = this->current;
	if (node->Left == NULL && node->Right == NULL) {
		switch (child) {
			case TREE44:
				node->Left = new TNode();
				this->nodeCount++;
				node->Left->Data = 10;
				node->Left->Parent = node;
				
				node->Right = new TNode();
				this->nodeCount++;
				node->Right->Data = 11;
				node->Right->Parent = node;
				break;
			case TREE45:
				if (node->Data % 2 != 0) {
					node->Left = new TNode();
					node->Left->Data = node->Data;
					node->Left->Parent = node;
				} else {
					node->Right = new TNode();
					node->Right->Data = node->Data;
					node->Right->Parent = node;
				}
				this->nodeCount++;
				break;
		}
		this->current = this->root;
		delete this->field;
		this->field = NULL;
		this->level = -1;
		return;
	}
	this->current = node->Left;
	this->addLeaves(child);
	this->current = node->Right;
	this->addLeaves(child);
}

void Tree::addSecondChild() {
	if (this->current == NULL) {
		this->current = this->root;
		delete this->field;
		this->field = NULL;
		this->level = -1;
		return;
	}
	TNode* node = this->current;
	if (node->Left == NULL && node->Right != NULL) {
		node->Left = new TNode();
		this->nodeCount++;
		node->Left->Data = node->Data * 2;
		node->Left->Parent = node;
	} else if (node->Right == NULL && node->Left != NULL) {
		node->Right = new TNode();
		this->nodeCount++;
		node->Right->Data = node->Data * 2;
		node->Right->Parent = node;
	}
	this->current = node->Left;
	this->addSecondChild();
	this->current = node->Right;
	this->addSecondChild();
}

void Tree::doPerfect(int level, int L) {
	if (level >= L) {
		this->current = this->root;
		delete this->field;
		this->field = NULL;
		return;
	}
	TNode* node = this->current;
	if (node->Left == NULL) {
		node->Left = new TNode();
		this->nodeCount++;
		node->Left->Data = -1;
		node->Left->Parent = node;
	}
	if (node->Right == NULL) {
		node->Right = new TNode();
		this->nodeCount++;
		node->Right->Data = -1;
		node->Right->Parent = node;
	}
	this->current = node->Left;
	this->doPerfect(level+1, L);
	this->current = node->Right;
	this->doPerfect(level+1, L);
}

TNode* Tree::getNode() {
	if (this->current == NULL) {
		this->current = this->root;
		return NULL;
	}
	TNode* node = this->current;
	cout << "current->Data = " << node->Data << ". Where to go? (0-exit; 1-left; 2-right; 3-current):\t";
	int answer;
	cin >> answer;
	switch (answer) {
		case 0: this->current = this->root; return NULL;
		case 1: this->current = node->Left; break;
		case 2: this->current = node->Right; break;
		case 3: this->current = this->root; return node;
	}
	return this->getNode();
}

TNode* Tree::getRoot(TNode* node) {
	if (node == NULL || node->Parent == NULL) {
		return node;
	}
	return this->getRoot(node->Parent);
}

void Tree::setParent() {
	if (this->current == NULL) {
		this->current = this->root;
		return;
	}
	TNode* node = this->current;
	if (node == this->root) {
		node->Parent = NULL;
	}
	if (node->Left != NULL) {
		node->Left->Parent = node;
		this->current = node->Left;
		this->setParent();
	}
	if (node->Right != NULL) {
		node->Right->Parent = node;
		this->current = node->Right;
		this->setParent();
	}
}

int Tree::getNodeHeight(TNode* node) {
	if (node == NULL) {
		return 0;
	}
	return 1 + this->getNodeHeight(node->Parent);
}

TNode* Tree::getClosestCommonParent(TNode* i, TNode* node1, TNode* j, TNode* node2) {
	if (i == NULL || i == j) {
		return i;
	}
	j = j == NULL ? i = i->Parent, node2 : j->Parent;
	return this->getClosestCommonParent(i, node1, j, node2);
}