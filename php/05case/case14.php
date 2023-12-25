<?php
$N = mt_rand(1, 4);
$qimat = mt_rand(1, 100) / 4;
$qimat = number_format($qimat, 2);
echo "N = ".$N;
echo "<br>Значение: ".$qimat;
switch($N){
	case 1:
		$a = $qimat;
		$R1 = $a * sqrt(3) / 6;
		$R2 = 2 * $R1;
		$S = $a * $a * sqrt(3) / 4;
		break;
	case 2:
		$R1 = $qimat;
		$a = 6 * $R1 / sqrt(3);
		$R2 = 2 * $R1;
		$S = $a * $a * sqrt(3) / 4;
		break;
	case 3:
		$R2 = $qimat;
		$R1 = $R2 / 2;
		$a = 6 * $R1 / sqrt(3);
		$S = $a * $a * sqrt(3) / 4;
		break;
	case 4:
		$S = $qimat;
		$a = sqrt(4 * $S / sqrt(3));
		$R1 = $a * sqrt(3) / 6;
		$R2 = 2 * $R1;
}
$S = number_format($S, 2);
$a = number_format($a, 2);
$R1 = number_format($R1, 2);
$R2 = number_format($R2, 2);
echo "<br>Площадь: ".$S;
echo "<br>Радиус вписанной окружности: ".$R1;
echo "<br>Радиус описанной окружности: ".$R2;
echo "<br>Сторона прямоугольного треугольника: ".$a;
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
		$R1 = $a * sqrt(3) / 6;
		$R2 = 2 * $R1;
		$S = $a * $a * sqrt(3) / 4;
		break;
	case 2:
		$R1 = $qimat;
		$a = 6 * $R1 / sqrt(3);
		$R2 = 2 * $R1;
		$S = $a * $a * sqrt(3) / 4;
		break;
	case 3:
		$R2 = $qimat;
		$R1 = $R2 / 2;
		$a = 6 * $R1 / sqrt(3);
		$S = $a * $a * sqrt(3) / 4;
		break;
	case 4:
		$S = $qimat;
		$a = sqrt(4 * $S / sqrt(3));
		$R1 = $a * sqrt(3) / 6;
		$R2 = 2 * $R1;
}
$S = number_format($S, 2);
$a = number_format($a, 2);
$R1 = number_format($R1, 2);
$R2 = number_format($R2, 2);
echo "<br>Масоҳат: ".$S;
echo "<br>Радиуси давраи дохилӣ: ".$R1;
echo "<br>Радиуси давраи берунӣ: ".$R2;
echo "<br>Тарафи секунҷаи росткунҷа: ".$a;
?>
*/