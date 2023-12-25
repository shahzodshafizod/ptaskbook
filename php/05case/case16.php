<?php
$N = mt_rand(20, 69);
echo "N = ".$N."<br>";
$a = intval($N / 10);
$b = $N % 10;
switch($a){
	case 2:
		echo "двадцать ";
		break;
	case 3:
		echo "тридцать ";
		break;
	case 4:
		echo "сорок ";
		break;
	case 5:
		echo "пятьдесят ";
		break;
	case 6:
		echo "шестьдесят ";
}
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
switch($b){
	case 1:
		echo "год";
		break;
	case 2: case 3: case 4:
		echo "года";
		break;
	default:
		echo "лет";
}
?>
