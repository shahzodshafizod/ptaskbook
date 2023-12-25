<?php
$N = mt_rand(1, 5);
$vazn = mt_rand(1, 1000000);
$vazn = number_format($vazn, 2);
echo "N = ".$N;
echo "<br>Вес = ".$vazn."<br>";
switch($N){
	case 1:
		echo "килограмм = ".$vazn;
		$kg = $vazn;
		break;
	case 2:
		echo "миллиграмм = ".$vazn;
		$kg = $vazn / 1000000;
		break;
	case 3:
		echo "грамм = ".$vazn;
		$kg = $vazn / 1000;
		break;
	case 4:
		echo "тонна = ".$vazn;
		$kg = $vazn * 1000;
		break;
	case 5:
		echo "центнер = ".$vazn;
		$kg = $vazn * 100;
}
echo "<br>килограмм = ".$kg;
?>

/*
<?php
$N = mt_rand(1, 5);
$vazn = mt_rand(1, 1000000);
$vazn = number_format($vazn, 2);
echo "N = ".$N;
echo "<br>вазн = ".$vazn."<br>";
switch($N){
	case 1:
		echo "килограм = ".$vazn;
		$kg = $vazn;
		break;
	case 2:
		echo "миллиграм = ".$vazn;
		$kg = $vazn / 1000000;
		break;
	case 3:
		echo "грам = ".$vazn;
		$kg = $vazn / 1000;
		break;
	case 4:
		echo "тонна = ".$vazn;
		$kg = $vazn * 1000;
		break;
	case 5:
		echo "сентнер = ".$vazn;
		$kg = $vazn * 100;
}
echo "<br>килограм = ".$kg;
?>
*/