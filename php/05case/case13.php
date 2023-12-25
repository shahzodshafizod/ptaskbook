<?php
$N = mt_rand(1, 4);
$qimat = mt_rand(1, 100) / 4;
$qimat = number_format($qimat, 2);
echo "N = ".$N;
echo "<br>Значение: ".$qimat;
switch($N){
	case 1:
		$a = $qimat;
		$c = $a * sqrt(2);
		$h = $c / 2;
		$S = $c * $h / 2;
		break;
	case 2:
		$c = $qimat;
		$a = $c / sqrt(2);
		$h = $c / 2;
		$S = $c * $h / 2;
		break;
	case 3:
		$h = $qimat;
		$c = 2 * $h;
		$a = c / sqrt(2);
		$S = $c * $h / 2;
		break;
	case 4:
		$S = $qimat;
		$h = sqrt($S);
		$c = 2 * $h;
		$a = $c / sqrt(2);
}
$S = number_format($S, 2);
$h = number_format($h, 2);
$c = number_format($c, 2);
$a = number_format($a, 2);
echo "<br>Площадь: ".$S;
echo "<br>Высота: ".$h;
echo "<br>Гипотенуза: ".$c;
echo "<br>Катет: ".$a;
?>

/*
<?php
$N = mt_rand(1, 4);
$qimat = mt_rand(1, 100) / 4;
$qimat = number_format($qimat, 2);
echo "N = ".$N;
echo "<br>Қимат: ".$qimat;
switch($N){
	case 1:
		$a = $qimat;
		$c = $a * sqrt(2);
		$h = $c / 2;
		$S = $c * $h / 2;
		break;
	case 2:
		$c = $qimat;
		$a = $c / sqrt(2);
		$h = $c / 2;
		$S = $c * $h / 2;
		break;
	case 3:
		$h = $qimat;
		$c = 2 * $h;
		$a = c / sqrt(2);
		$S = $c * $h / 2;
		break;
	case 4:
		$S = $qimat;
		$h = sqrt($S);
		$c = 2 * $h;
		$a = $c / sqrt(2);
}
$S = number_format($S, 2);
$h = number_format($h, 2);
$c = number_format($c, 2);
$a = number_format($a, 2);
echo "<br>Масоҳат: ".$S;
echo "<br>Баландӣ: ".$h;
echo "<br>Гипотенуза: ".$c;
echo "<br>Катет: ".$a;
?>
*/