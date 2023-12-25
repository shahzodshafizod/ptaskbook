<?php
$M = mt_rand(1, 12);
do{
	$D = mt_rand(1, 31);
}
while(($M==2)&&($D>28) || ($M==4)&&($M==6)&&($M==9)&&($M==11)&&($D>30));
echo "Указанная дата:<br>";
echo $D.". ".$M;
switch($M){
	case 1: case 3:
	case 5: case 7:
	case 8: case 10:
		if($D == 31){
			$D = 1;
			$M++;
		}
		else
			$D++;
		break;
	case 4: case 6:
	case 9: case 11:
		if($D == 30){
			$D = 1;
			$M++;
		}
		else
			$D++;
		break;
	case 2:
		if($D == 28){
			$D = 1;
			$M++;
		}
		else
			$D++;
		break;
	case 12:
		if($D == 31){
			$D = 1;
			$M = 1;
		}
		else
			$D++;
}
echo "<br>Дата, следующая за указанной:<br>";
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
	case 1: case 3:
	case 5: case 7:
	case 8: case 10:
		if($D == 31){
			$D = 1;
			$M++;
		}
		else
			$D++;
		break;
	case 4: case 6:
	case 9: case 11:
		if($D == 30){
			$D = 1;
			$M++;
		}
		else
			$D++;
		break;
	case 2:
		if($D == 28){
			$D = 1;
			$M++;
		}
		else
			$D++;
		break;
	case 12:
		if($D == 31){
			$D = 1;
			$M = 1;
		}
		else
			$D++;
}
echo "<br>Cанаи як рӯз пас:<br>";
echo $D.". ".$M;
?>
*/