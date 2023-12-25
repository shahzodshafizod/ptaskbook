<?php

require_once "Flags.php";
require_once "Field.php";

class Tree {
	private $root;
	private $current;
	private $field;
	private $nodeCount;
	private $level;
	public function __construct() {
		$this->root = null;
		$this->current = null;
		$this->field = null;
		$this->nodeCount = 0;
		$this->level = -1;
	}
	public function automake($count) {
		if ($count <= 0) {
			$this->current = $this->root;
			return;
		}
		if ($this->root == null) {
			$this->root = new Node(rand(1, 100));
			$this->nodeCount++;
			$count--;
			$this->current = $this->root;
			if ($count <= 0) {
				return;
			}
		}
		$direction = 0;
		do {
			$direction = rand(1, 3);
		} while ($direction == 3 && $this->current->parent == null);
		$exit = false;
		switch ($direction) {
			case 1:
				if ($this->current->left != null) {
					$this->current = $this->current->left;
					$exit = true;
				}
				break;
			case 2:
				if ($this->current->right != null) {
					$this->current = $this->current->right;
					$exit = true;
				}
				break;
			case 3:
				$this->current = $this->current->parent;
				$exit = true;
				break;
		}
		if ($exit) {
			$this->automake($count);
			return;
		}
		$newNode = new Node(rand(1, 100));
		$this->nodeCount++;
		$count--;
		$newNode->parent = $this->current;
		switch ($direction) {
			case 1: $this->current->left = $newNode; break;
			case 2: $this->current->right = $newNode; break;
		}
		$this->current = $newNode;
		$this->automake($count);
	}
	public function print($flag) {
		if ($this->current == null) {
			$this->current = $this->root;
			return;
		}
		$node = $this->current;
		if ($flag == PREFIX) { echo $node->data.", "; }
		$this->current = $node->left;
		$this->print($flag);
		if ($flag == INFIX) { echo $node->data.", "; }
		$this->current = $node->right;
		$this->print($flag);
		if ($flag == POSTFIX) { echo $node->data.", "; }
	}
	public function getNodeCount() {
		return $this->nodeCount;
	}
	public function getLevel() {
		if ($this->level < 0) {
			$this->setLevel(0);
		}
		return $this->level;
	}
	private function setLevel($level) {
		if ($this->current == null) {
			$this->current = $this->root;
			return;
		}
		if ($level > $this->level) {
			$this->level = $level;
		}
		$node = $this->current;
		$this->current = $node->left;
		$this->setLevel($level+1);
		$this->current = $node->right;
		$this->setLevel($level+1);
	}
	public function display() {
		if ($this->root == null) { return; }
		if ($this->field == null) {
			$this->field = new Field($this->root, $this->getLevel()+1, $this->getNodeCount());
		}
		$this->field->display();
	}
}

?>