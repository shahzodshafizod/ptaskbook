<?php
const PI = 3.14;
$N = mt_rand(1, 4);
$qimat = mt_rand(1, 100) / 4;
$qimat = number_format($qimat, 2);
echo "N = ".$N;
echo "<br>Значение: ".$qimat;
switch($N){
	case 1:
		$R = $qimat;
		$D = 2 * $R;
		$L = 2 * PI * $R;
		$S = PI * $R * $R;
		break;
	case 2:
		$D = $qimat;
		$R = $D / 2;
		$L = 2 * PI * $R;
		$S = PI * $R * $R;
		break;
	case 3:
		$L = $qimat;
		$R = $L / (2 * PI);
		$D = 2 * $R;
		$S = PI * $R * $R;
		break;
	case 4:
		$S = $qimat;
		$R = sqrt($S / PI);
		$D = 2 * $R;
		$L = 2 * PI * $R;
}
$S = number_format($S, 2);
$R = number_format($R, 2);
$D = number_format($D, 2);
$L = number_format($L, 2);
echo "<br>Площадь круга: ".$S;
echo "<br>Радиус: ".$R;
echo "<br>Диаметр: ".$D;
echo "<br>Длина: ".$L;
?>

/*
<?php
const PI = 3.14;
$N = mt_rand(1, 4);
$qimat = mt_rand(1, 100) / 4;
$qimat = number_format($qimat, 2);
echo "N = ".$N;
echo "<br>Қимат: ".$qimat;
switch($N){
	case 1:
		$R = $qimat;
		$D = 2 * $R;
		$L = 2 * PI * $R;
		$S = PI * $R * $R;
		break;
	case 2:
		$D = $qimat;
		$R = $D / 2;
		$L = 2 * PI * $R;
		$S = PI * $R * $R;
		break;
	case 3:
		$L = $qimat;
		$R = $L / (2 * PI);
		$D = 2 * $R;
		$S = PI * $R * $R;
		break;
	case 4:
		$S = $qimat;
		$R = sqrt($S / PI);
		$D = 2 * $R;
		$L = 2 * PI * $R;
}
$S = number_format($S, 2);
$R = number_format($R, 2);
$D = number_format($D, 2);
$L = number_format($L, 2);
echo "<br>Масоҳати доира: ".$S;
echo "<br>Радиус: ".$R;
echo "<br>Диаметр: ".$D;
echo "<br>Дарозӣ: ".$L;
?>
*/