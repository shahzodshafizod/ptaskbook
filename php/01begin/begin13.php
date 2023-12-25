<?php
const PI = 3.14;
do{
	$R1 = mt_rand(1, 100) / 4;
	$R2 = mt_rand(1, 100) / 4;
}
while($R1 <= $R2);
$S1 = PI * $R1 * $R1;
$S2 = PI * $R2 * $R2;
$S3 = PI * $R3 * $R3;
$R1 = number_format($R1, 2);
$R2 = number_format($R2, 2);
$S1 = number_format($S1, 2);
$S2 = number_format($S2, 2);
$S3 = number_format($S3, 2);
echo "Внешний радиус: ".$R1;
echo "<br>Внутренний радиус: ".$R2;
echo "<br>Площадь первого круга: ".$S1;
echo "<br>Площадь второго круга: ".$S2;
echo "<br>Площадь кольца: ".$S3;
?>

/*
<?php
const PI = 3.14;
do{
	$R1 = mt_rand(1, 100) / 4;
	$R2 = mt_rand(1, 100) / 4;
}
while($R1 <= $R2);
$S1 = PI * $R1 * $R1;
$S2 = PI * $R2 * $R2;
$S3 = PI * $R3 * $R3;
$R1 = number_format($R1, 2);
$R2 = number_format($R2, 2);
$S1 = number_format($S1, 2);
$S2 = number_format($S2, 2);
$S3 = number_format($S3, 2);
echo "Радиуси беруна: ".$R1;
echo "<br>Радиуси дохилӣ: ".$R2;
echo "<br>Масоҳати доираи якӯм: ".$S1;
echo "<br>Масоҳати доираи дуюм: ".$S2;
echo "<br>Масоҳати ҳалқа: ".$S3;
?>
*/