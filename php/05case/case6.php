<?php
$N = mt_rand(1, 5);
$darozi = mt_rand(1, 1000000) / 4;
$N = number_format($N, 2);
echo "N = ".$N;
echo "<br>Длина = ".$darozi."<br>";
switch($N){
	case 1:
		echo "дм = ".$darozi;
		$meter = $darozi / 10;
		break;
	case 2:
		echo "км = ".$darozi;
		$meter = $darozi * 1000;
		break;
	case 3:
		echo "м = ".$darozi;
		$meter = $darozi;
		break;
	case 4:
		echo "мм = ".$darozi;
		$meter = $darozi / 1000;
		break;
	case 5:
		echo "см = ".$darozi;
		$meter = $darozi / 100;
}
$meter = number_format($meter, 2);
echo "<br>метр = ".$meter;
?>

/*
<?php
$N = mt_rand(1, 5);
$darozi = mt_rand(1, 1000000) / 4;
$N = number_format($N, 2);
echo "N = ".$N;
echo "<br>дарозӣ = ".$darozi."<br>";
switch($N){
	case 1:
		echo "дм = ".$darozi;
		$meter = $darozi / 10;
		break;
	case 2:
		echo "км = ".$darozi;
		$meter = $darozi * 1000;
		break;
	case 3:
		echo "м = ".$darozi;
		$meter = $darozi;
		break;
	case 4:
		echo "мм = ".$darozi;
		$meter = $darozi / 1000;
		break;
	case 5:
		echo "см = ".$darozi;
		$meter = $darozi / 100;
}
$meter = number_format($meter, 2);
echo "<br>метр = ".$meter;
?>
*/