<?php
$N = mt_rand(100, 999);
$a = intval($N / 100);
$b = intval($N / 10) % 10;
$c = $N % 10;
echo "N = ".$N."<br>";
switch($a){
	case 1:
		echo "сто ";
		break;
	case 2:
		echo "двести ";
		break;
	case 3:
		echo "триста ";
		break;
	case 4:
		echo "четыреста ";
		break;
	case 5:
		echo "пятьсот ";
		break;
	case 6:
		echo "шестьсот ";
		break;
	case 7:
		echo "семьсот ";
		break;
	case 8:
		echo "восемьсот ";
		break;
	case 9:
		echo "девятьсот ";
}
switch($b){
	case 1:
		switch($c){
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
		break;
	case 5:
		echo "пятьдесят ";
		break;
	case 6:
		echo "шестьдесят ";
		break;
	case 7:
		echo "семьдесят ";
		break;
	case 8:
		echo "восемьдесят ";
		break;
	case 9:
		echo "девяносто ";
}
if($b != 1){
	switch($c){
		case 1:
			echo "один";
			break;
		case 2:
			echo "два";
			break;
		case 3:
			echo "три";
			break;
		case 4:
			echo "четыре";
			break;
		case 5:
			echo "пять";
			break;
		case 6:
			echo "шесть";
			break;
		case 7:
			echo "семь";
			break;
		case 8:
			echo "восемь";
			break;
		case 9:
			echo "девять";
	}
}
?>

/*
<?php
$N = mt_rand(100, 999);
$a = intval($N / 100);
$b = intval($N / 10) % 10;
$c = $N % 10;
echo "N = ".$N."<br>";
switch($a){
	case 1:
		echo "сад";
		break;
	case 2:
		echo "дусад";
		break;
	case 3:
		echo "сесад";
		break;
	case 4:
		echo "чорсад";
		break;
	case 5:
		echo "панҷсад";
		break;
	case 6:
		echo "шашсад";
		break;
	case 7:
		echo "ҳафтсад";
		break;
	case 8:
		echo "ҳаштсад";
		break;
	case 9:
		echo "нӯҳсад";
}
if(($b != 0) && ($c != 0))
		echo "у ";
switch($b){
	case 1:
		switch($c){
			case 0:
				echo "даҳ";
				break;
			case 1:
				echo "ёздаҳ";
				break;
			case 2:
				echo "дувоздаҳ";
				break;
			case 3:
				echo "сездаҳ";
				break;
			case 4:
				echo "чордаҳ";
				break;
			case 5:
				echo "понздаҳ";
				break;
			case 6:
				echo "шонздаҳ";
				break;
			case 7:
				echo "ҳабдаҳ";
				break;
			case 8:
				echo "ҳаждаҳ";
				break;
			case 9:
				echo "нуздаҳ";
		}
		break;
	case 2:
		echo "бист";
		break;
	case 3:
		echo "сӣ";
		break;
	case 4:
		echo "чил";
		break;
	case 5:
		echo "панҷоҳ";
		break;
	case 6:
		echo "шаст";
		break;
	case 7:
		echo "ҳафтод";
		break;
	case 8:
		echo "ҳаштод";
		break;
	case 9:
		echo "навад";
}
if($b == 3)
	echo "ю ";
else if(($c != 0) && ($b != 1))
	echo "у ";
if($b != 1){
	switch($c){
		case 1:
			echo "як";
			break;
		case 2:
			echo "ду";
			break;
		case 3:
			echo "се";
			break;
		case 4:
			echo "чор";
			break;
		case 5:
			echo "панҷ";
			break;
		case 6:
			echo "шаш";
			break;
		case 7:
			echo "ҳафт";
			break;
		case 8:
			echo "ҳашт";
			break;
		case 9:
			echo "нӯҳ";
	}
}
?>
*/