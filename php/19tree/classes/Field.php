<?php

require_once "Node.php";

class Field {
	private $nodes;
	private $chars;
	private $heigth;
	private $width;
	private $left;
	private $center;
	private $right;
	private $cellWidth;
	public function __construct($node, $height, $width) {
		$this->height = $height;
		$this->width = $width;
		$this->cellWidth = 4;
		$this->left = '<';
		$this->right = '>';
		$this->center = '-';
		$this->init($node);
	}
	private function init($node) {
		for ($i = 0; $i < $this->height; $i++) {
			for ($j = 0; $j < $this->width; $j++) {
				$this->nodes[$i][$j] = null;
				$this->chars[$i][$j] = SPACE;
			}
		}
		$col = 0;
		$this->fillNodes($node, 0, $col);
		$this->fillChars();
	}
	private function fillNodes($node, $row, &$col) {
		if ($node == null) {
			return;
		}
		$this->fillNodes($node->left, $row+1, $col);
		$this->nodes[$row][$col++] = $node;
		$this->fillNodes($node->right, $row+1, $col);
	}
	private function fillChars() {
		for ($i = 0; $i < $this->height; $i++) {
			for ($j = 0; $j < $this->width; $j++) {
				if ($this->nodes[$i][$j] == null) {
					continue;
				}
				$this->chars[$i][$j] = DATA;
				if ($this->nodes[$i][$j]->left != null) {
					for ($k = $j-1; $k >= 0; $k--) {
						if ($this->nodes[$i+1][$k] != null) {
							$this->chars[$i][$k] = LEFT;
							break;
						}
						$this->chars[$i][$k] = CENTER;
					}
				}
				if ($this->nodes[$i][$j]->right != null) {
					for ($j++; $j < $this->width; $j++) {
						if ($this->nodes[$i+1][$j] != null) {
							$this->chars[$j][$j] = RIGHT;
						}
						$this->chars[$i][$j] = CENTER;
					}
				}
			}
		}
	}
	public function display() {
		for ($i = 0; $i < $this->height; $i++) {
			$nextNode = null;
			$data = 0;
			$spaces = 0;
			$half = 0;
			for ($j = 0; $j < $this->width; $j++) {
				switch ($this->chars[$i][$j]) {
					case SPACE:
						$this->printChars(' ', $this->cellWidth);
						break;
					case LEFT:
						$nextNode = $this->getNode($i, $j);
						$data = $nextNode->data;
						$spaces = $this->cellWidth - $this->getLength($data);
						$half = intval($spaces / 2);
						$this->printChars(' ', $half);
						$this->printChars($this->left, 1);
						$this->printChars($this->center, $this->cellWidth - 1);
						break;
					case CENTER:
						$this->printChars($this->center, $this->cellWidth);
						break;
					case RIGHT:
						$this->printChars($this->center, $this->cellWidth - 1);
						$this->printChars($this->right, 1);
						$this->printChars(' ', $spaces - $half);
						break;
					case DATA:
						if ($this->nodes[$i][$j]->left == null) {
							$nextNode = $this->nodes[$i][$j];
							$data = $nextNode->data;
							$spaces = $this->cellWidth - $this->getLength($data);
							$half = intval($spaces / 2);
							$this->printChars(' ', $half);
						}
						echo $data;
						if ($this->nodes[$i][$j]->right == null) {
							$this->printChars(' ', $spaces - $half);
						}
						break;
				}
			}
			echo "\n";
		}
	}
	private function printChars($c, $count) {
		for ($i = 0; $i < $count; $i++) {
			echo $c;
		}
	}
	private function getLength($data) {
		$len = 0;
		if ($data < 0) {
			$len++;
			$data *= -1;
		}
		while ($data > 0) {
			$len++;
			$data = intval($data / 10);
		}
		return $len;
	}
	private function getNode($row, $col) {
		for ($col++; $col < $this->width; $col++) {
			if ($this->nodes[$row][$col] != null) {
				return $this->nodes[$row][$col];
			}
		}
		return null;
	}
}

?>