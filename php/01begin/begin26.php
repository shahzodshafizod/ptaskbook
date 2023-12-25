<?php
$x = mt_rand(-100, 100) / 4;
$y = 4 * pow($x - 3, 6) - 7 * pow($x - 3, 3) + 2;
$x = number_format($x, 2);
$y = number_format($y, 2);
echo "Аргумент: ".$x;
echo "<br>Функция: ".$y;
?>

/*
<?php
$x = mt_rand(-100, 100) / 4;
$y = 4 * pow($x - 3, 6) - 7 * pow($x - 3, 3) + 2;
$x = number_format($x, 2);
$y = number_format($y, 2);
echo "Аргумент: ".$x;
echo "<br>Функсия: ".$y;
?>
*/