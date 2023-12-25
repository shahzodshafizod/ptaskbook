#ifndef ____TREE____H____
#define ____TREE____H____
#include "TNode.h"
#include "Field.h"
#include <cstdlib>

enum ChildNode { TREE25, TREE26LEFT, TREE26RIGHT, TREE27, TREE28, TREE29, 
	TREE31LEFT, TREE31RIGHT, TREE35, TREE36, TREE37, TREE38, TREE42, TREE43,
	TREE44, TREE45 };

class Tree {
	private: TNode* root;
	private: TNode* current;
	private: int nodeCount;
	private: int level;
	private: Field* field;
	
	public: Tree(); //Contructor
	public: Tree(Tree&); //Copy contructor
	public: ~Tree(); //Destructor
	public: void make();
	public: void make(int[], int, int, ChildNode); //Tree 25, 26, 27, 28, 29
	public: void make(int[], int&, int, int, int, bool toLeft = true); //Tree31
	public: void make(int[], int&, int, int, bool toLeft = true); //Tree32
	public: void make(int, int, bool toLeft = true); //Tree33
	public: void make(int); //Tree 30
	public: void copy(Tree&, bool toRight = false); //Tree34
	public: void copy(TNode*, bool toRight = false); //Tree54
	private: void free();
	public: void display(TNode* node = NULL);
	public: void infix(); //Tree 12
	public: void prefix(); //Tree 13
	public: void postfix(); //Tree 14
	public: void infixToN(int&, int); //Tree 15
	public: void postfixFromN(int&, int); //Tree 16
	public: void prefixBetween(int&, int, int); //Tree 17
	public: int getNodeCount() const; //Tree 2
	public: int getLeftCount(bool isLeft = false); //Tree 5
	public: int getLeafCount(); //Tree 6
	public: int getRightLeafCount(bool isRight = false); // Tree 8
	public: int getNodeCountK(int); //Tree 3
	public: int getLeafCountK(int); //Tree 20
	private: void setLevel(int currentLevel = 0);
	public: int getLevel(); //Tree 9
	public: int getLevelNodeCount(int, int level = 0); //Tree 18
	public: int getDataSum(); //Tree 4
	public: int getLeafDataSum(); //Tree 7
	public: void levelNodeCountToArr(int[], int); //Tree 10
	public: void levelNodeSumToArr(int[], int); //Tree 11
	public: int getMaxData(); //Tree 19
	public: int getMinData(); //Tree 20
	public: int getMinLeafData(); //Tree 21
	public: int getMaxInternalData(); //Tree 22
	public: TNode* getFirstNodePrefix(int); //Tree 23
	public: TNode* getLastNodeInfix(int); //Tree 24
	public: bool hasOddData(); //Tree 24
	public: int getMaxOddData(); //Tree 24
	public: void changeDatas(ChildNode); //Tree35, Tree36, Tree37, Tree38
	public: void swapChilds(); //Tree39
	public: void deleteAllChilds(); //Tree40
	public: void deleteAllLeaves(); //Tree41
	public: void deleteSubTree(ChildNode); //Tree42, Tree43
	public: void addLeaves(ChildNode); //Tree44, Tree45
	public: void addSecondChild(); //Tree46
	public: void doPerfect(int, int); //Tree47
	public: TNode* getNode(); //Tree48
	public: TNode* getRoot(TNode*); //Tree50
	public: void setParent(); //Tree49
	public: int getNodeHeight(TNode*); //Tree51
	public: TNode* getClosestCommonParent(TNode*, TNode*, TNode*, TNode*); //Tree53
};

#endif