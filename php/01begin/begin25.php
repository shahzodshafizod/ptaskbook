<?php
$x = mt_rand(-100, 100) / 4;
$y = 3 * pow($x, 6) - 6 * $x * $x - 7;
$x = number_format($x, 2);
$y = number_format($y, 2);
echo "Аргумент: ".$x;
echo "<br>Функция: ".$y;
?>

/*
<?php
$x = mt_rand(-100, 100) / 4;
$y = 3 * pow($x, 6) - 6 * $x * $x - 7;
$x = number_format($x, 2);
$y = number_format($y, 2);
echo "Аргумент: ".$x;
echo "<br>Функсия: ".$y;
?>
*/