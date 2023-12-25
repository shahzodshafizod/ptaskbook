<?php
$M = mt_rand(1, 12);
do{
	$D = mt_rand(1, 31);
}
while(($M==2)&&($D>28) || ($M==4)&&($M==6)&&($M==9)&&($M==11)&&($D>30));
echo "Указанная дата:<br>";
echo $D.". ".$M;
switch($M){
	case 4: case 6:
	case 8: case 9:
	case 2: case 11:
		if($D == 1){
			$D = 31;
			$M--;
		}
		else
			$D--;
		break;
	case 5: case 7:
	case 10: case 12:
		if($D == 1){
			$D = 30;
			$M--;
		}
		else
			$D--;
		break;
	case 1:
		if($D == 1){
			$D = 31;
			$M = 12;
		}
		else
			$D--;
		break;
	case 3:
		if($D == 1){
			$D = 28;
			$M--;
		}
		else
			$D--;
}
echo "<br>Дата, предшествующая указанной:<br>";
echo $D.". ".$M;
?>

/*
<?php
$M = mt_rand(1, 12);
do{
	$D = mt_rand(1, 31);
}
while(($M==2)&&($D>28) || ($M==4)&&($M==6)&&($M==9)&&($M==11)&&($D>30));
echo "Санаи додашуда:<br>";
echo $D.". ".$M;
switch($M){
	case 4: case 6:
	case 8: case 9:
	case 2: case 11:
		if($D == 1){
			$D = 31;
			$M--;
		}
		else
			$D--;
		break;
	case 5: case 7:
	case 10: case 12:
		if($D == 1){
			$D = 30;
			$M--;
		}
		else
			$D--;
		break;
	case 1:
		if($D == 1){
			$D = 31;
			$M = 12;
		}
		else
			$D--;
		break;
	case 3:
		if($D == 1){
			$D = 28;
			$M--;
		}
		else
			$D--;
}
echo "<br>Cанаи як рӯз пеш:<br>";
echo $D.". ".$M;
?>
*/