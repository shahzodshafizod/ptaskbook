<?php
const PI = 3.14;
$d = mt_rand(1, 50) / 4;
$L = PI * $d;
$d = number_format($d, 2);
$L = number_format($L, 2);
echo "Диаметр окружности: ".$d;
echo "<br>Длина окружности: ".$L;
?>

/*
<?php
const PI = 3.14;
$d = mt_rand(1, 50) / 4;
$L = PI * $d;
$d = number_format($d, 2);
$L = number_format($L, 2);
echo "Диаметри давра: ".$d;
echo "<br>Дарозии давра: ".$L;
?>
*/