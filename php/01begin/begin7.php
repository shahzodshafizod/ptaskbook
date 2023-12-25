<?php
const PI = 3.14;
$R = mt_rand(1, 50) / 4;
$L = 2 * PI * $R;
$S = PI * $R * $R;
$R = number_format($R, 2);
$L = number_format($L, 2);
$S = number_format($S, 2);
echo "Радиус окружности: ".$R;
echo "<br>Длина окружности: ".$L;
echo "<br>Площадь круга: ".$S;
?>

/*
<?php
const PI = 3.14;
$R = mt_rand(1, 50) / 4;
$L = 2 * PI * $R;
$S = PI * $R * $R;
$R = number_format($R, 2);
$L = number_format($L, 2);
$S = number_format($S, 2);
echo "Радиуси давра: ".$R;
echo "<br>Дарозии давра: ".$L;
echo "<br>Масоҳати доира: ".$S;
?>
*/