<?php
const PI = 3.14;
$S = mt_rand(1, 100) / 4;
$D = sqrt(4 * $S / PI);
$L = PI * $D;
$S = number_format($S, 2);
$L = number_format($L, 2);
$D = number_format($D, 2);
echo "Площадь круга: ".$S;
echo "<br>Диаметр круга: ".$D;
echo "<br>Длина окружности: ".$L;
?>

/*
<?php
const PI = 3.14;
$S = mt_rand(1, 100) / 4;
$D = sqrt(4 * $S / PI);
$L = PI * $D;
$S = number_format($S, 2);
$L = number_format($L, 2);
$D = number_format($D, 2);
echo "Масоҳати доира: ".$S;
echo "<br>Диаметри доира: ".$D;
echo "<br>Дарозии давра: ".$L;
?>
*/