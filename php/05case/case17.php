<?php
$N = mt_rand(10, 40);
echo "N = ".$N."<br>";
$a = intval($N / 10);
$b = $N % 10;
switch($a){
	case 1:
		switch($b){
			case 0:
				echo "десять ";
				break;
			case 1:
				echo "одиннадцать ";
				break;
			case 2:
				echo "двенадцать ";
				break;
			case 3:
				echo "тринадцать ";
				break;
			case 4:
				echo "четырнадцать ";
				break;
			case 5:
				echo "пятнадцать ";
				break;
			case 6:
				echo "шестнадцать ";
				break;
			case 7:
				echo "семнадцать ";
				break;
			case 8:
				echo "восемнадцать ";
				break;
			case 9:
				echo "девятнадцать ";
		}
		break;
	case 2:
		echo "двадцать ";
		break;
	case 3:
		echo "тридцать ";
		break;
	case 4:
		echo "сорок ";
}
if($a != 1){
	switch($b){
		case 0:
			echo "";
			break;
		case 1:
			echo "один ";
			break;
		case 2:
			echo "два ";
			break;
		case 3:
			echo "три ";
			break;
		case 4:
			echo "четыре ";
			break;
		case 5:
			echo "пять ";
			break;
		case 6:
			echo "шесть ";
			break;
		case 7:
			echo "семь ";
			break;
		case 8:
			echo "восемь ";
			break;
		case 9:
			echo "девять ";
	}
}
if($a != 1)
	switch($b){
		case 1:
			echo "учебное задание";
			break;
		case 2: case 3: case 4:
			echo "учебного задания";
			break;
		default:
			echo "учебных заданий";
	}
else
	echo "учебных заданий";
?>
