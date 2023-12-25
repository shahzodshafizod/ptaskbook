<?php
$M = mt_rand(1, 12);
do{
	$D = mt_rand(1, 31);
}
while(($M==2)&&($D>28) || ($M==4)&&($M==6)&&($M==9)&&($M==11)&&($D>30));
echo "день: ".$D."<br>";
echo "месяц: ".$M."<br>";
echo "Знак Зодиака: ";
switch($M){
	case 1:
		if($D <= 19)
			echo "Козерог";
		else
			echo "Водолей";
		break;
	case 2:
		if($D <= 18)
			echo "Водолей";
		else
			echo "Рыбы";
		break;
	case 3:
		if($D <= 20)
			echo "Рыбы";
		else
			echo "Овен";
		break;
	case 4:
		if($D <= 19)
			echo "Овен";
		else
			echo "Телец";
		break;
	case 5:
		if($D <= 20)
			echo "Телец";
		else
			echo "Близнецы";
		break;
	case 6:
		if($D <= 21)
			echo "Близнецы";
		else
			echo "Рак";
		break;
	case 7:
		if($D <= 22)
			echo "Рак";
		else
			echo "Лев";
		break;
	case 8:
		if($D <= 22)
			echo "Лев";
		else
			echo "Дева";
		break;
	case 9:
		if($D <= 22)
			echo "Дева";
		else
			echo "Весы";
		break;
	case 10:
		if($D <= 22)
			echo "Весы";
		else
			echo "Скорпион";
		break;
	case 11:
		if($D <= 11)
			echo "Скорпион";
		else
			echo "Стрелец";
		break;
	case 12:
		if($D <= 21)
			echo "Стрелец";
		else
			echo "Козерог";
}
?>

/*
<?php
$M = mt_rand(1, 12);
do{
	$D = mt_rand(1, 31);
}
while(($M==2)&&($D>28) || ($M==4)&&($M==6)&&($M==9)&&($M==11)&&($D>30));
echo "рӯз: ".$D."<br>";
echo "моҳ: ".$M."<br>";
echo "Бурҷи ";
switch($M){
	case 1:
		if($D <= 19)
			echo "Ҷадӣ";
		else
			echo "Давл";
		break;
	case 2:
		if($D <= 18)
			echo "Давл";
		else
			echo "Ҳут";
		break;
	case 3:
		if($D <= 20)
			echo "Ҳут";
		else
			echo "Ҳамал";
		break;
	case 4:
		if($D <= 19)
			echo "Ҳамал";
		else
			echo "Савр";
		break;
	case 5:
		if($D <= 20)
			echo "Савр";
		else
			echo "Ҷавзо";
		break;
	case 6:
		if($D <= 21)
			echo "Ҷавзо";
		else
			echo "Саратон";
		break;
	case 7:
		if($D <= 22)
			echo "Саратон";
		else
			echo "Асад";
		break;
	case 8:
		if($D <= 22)
			echo "Асад";
		else
			echo "Сунбула";
		break;
	case 9:
		if($D <= 22)
			echo "Сунбула";
		else
			echo "Мизон";
		break;
	case 10:
		if($D <= 22)
			echo "Мизон";
		else
			echo "Ақраб";
		break;
	case 11:
		if($D <= 11)
			echo "Ақраб";
		else
			echo "Қавс";
		break;
	case 12:
		if($D <= 21)
			echo "Қавс";
		else
			echo "Ҷадӣ";
}
?>
*/