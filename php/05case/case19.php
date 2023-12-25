<?php
$sol = mt_rand(1, 10000);
echo $sol."-ый год - год ";
$n = $sol % 60;
$m = $sol % 60 % 12;
if(($m == 6) || ($m == 7) || ($m == 8)){
	if(($n >= 4) && ($n < 16))
		echo "зеленого ";
	else if(($n >= 16) && ($n < 28))
		echo "красного ";
	else if(($n >= 28) && ($n < 40))
		echo "желтого ";
	else if(($n >= 40) && ($n < 52))
		echo "белого ";
	else if(($n >= 52) || ($n < 4))
		echo "черного ";
}
else{
	if(($n >= 4) && ($n < 16))
		echo "зеленой ";
	else if(($n >= 16) && ($n < 28))
		echo "красной ";
	else if(($n >= 28) && ($n < 40))
		echo "желтой ";
	else if(($n >= 40) && ($n < 52))
		echo "белой ";
	else if(($n >= 52) || ($n < 4))
		echo "черной ";
}
$n = $sol % 60 % 12;
switch($n){
	case 0:
		echo "обезяны";
		break;
	case 1:
		echo "курицы";
		break;
	case 2:
		echo "собаки";
		break;
	case 3:
		echo "свиньи";
		break;
	case 4:
		echo "крисы";
		break;
	case 5:
		echo "коровы";
		break;
	case 6:
		echo "тигра";
		break;
	case 7:
		echo "зайца";
		break;
	case 8:
		echo "дракона";
		break;
	case 9:
		echo "змеи";
		break;
	case 10:
		echo "лощади";
		break;
	case 11:
		echo "овцы";
}
?>

/*
<?php
$sol = mt_rand(1, 10000);
echo "Соли ".$sol." - соли ";
$n = $sol % 60 % 12;
switch($n){
	case 0:
		echo "маймуни ";
		break;
	case 1:
		echo "мурғи ";
		break;
	case 2:
		echo "саги ";
		break;
	case 3:
		echo "хуки ";
		break;
	case 4:
		echo "муши ";
		break;
	case 5:
		echo "гови ";
		break;
	case 6:
		echo "паланги ";
		break;
	case 7:
		echo "харгӯши ";
		break;
	case 8:
		echo "аждаҳори ";
		break;
	case 9:
		echo "мори ";
		break;
	case 10:
		echo "аспи ";
		break;
	case 11:
		echo "гӯспанди ";
}
$n = $sol % 60;
if(($n >= 4) && ($n < 16))
	echo "сабз";
else if(($n >= 16) && ($n < 28))
	echo "сурх";
else if(($n >= 28) && ($n < 40))
	echo "зард";
else if(($n >= 40) && ($n < 52))
	echo "сафед";
else if(($n >= 52) || ($n < 4))
	echo "сиёҳ";
echo " аст.";
?>
*/