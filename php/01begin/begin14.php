<?php
const PI = 3.14;
$L = mt_rand(1, 100) / 4;
$R = $L / (2 * PI);
$S = PI * $R * $R;
$R = number_format($R, 2);
$S = number_format($S, 2);
$L = number_format($L, 2);
echo "Длина окружности: ".$L;
echo "<br>Радиус окружности: ".$R;
echo "<br>Площадь круга: ".$S;
?>

/*
<?php
const PI = 3.14;
$L = mt_rand(1, 100) / 4;
$R = $L / (2 * PI);
$S = PI * $R * $R;
$R = number_format($R, 2);
$S = number_format($S, 2);
$L = number_format($L, 2);
echo "Дарозии давра: ".$L;
echo "<br>Радиуси давра: ".$R;
echo "<br>Масоҳати доира: ".$S;
?>
*/