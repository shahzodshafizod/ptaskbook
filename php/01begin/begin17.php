<?php
$A = mt_rand(-100, 100) / 4;
$B = mt_rand(-100, 100) / 4;
$C = mt_rand(-100, 100) / 4;
$AC = abs($C - $A);
$BC = abs($C - $B);
$jamm = $AC + $BC;
$A = number_format($A, 2);
$B = number_format($B, 2);
$C = number_format($C, 2);
$AC = number_format($AC, 2);
$BC = number_format($BC, 2);
$jamm = number_format($jamm, 2);
echo "Первая точка: ".$A;
echo "<br>Вторая точка: ".$B;
echo "<br>Третья точка: ".$C;
echo "<br>Длина отрезка AC: ".$AC;
echo "<br>Длина отрезка BC: ".$BC;
echo "<br>Сумма отрезков AC и BC: ".$jamm;
?>

/*
<?php
$A = mt_rand(-100, 100) / 4;
$B = mt_rand(-100, 100) / 4;
$C = mt_rand(-100, 100) / 4;
$AC = abs($C - $A);
$BC = abs($C - $B);
$jamm = $AC + $BC;
$A = number_format($A, 2);
$B = number_format($B, 2);
$C = number_format($C, 2);
$AC = number_format($AC, 2);
$BC = number_format($BC, 2);
$jamm = number_format($jamm, 2);
echo "Нуқтаи якӯм: ".$A;
echo "<br>Нуқтаи дуюм: ".$B;
echo "<br>Нуқтаи сеюм: ".$C;
echo "<br>Дарозии порчаи AC: ".$AC;
echo "<br>Дарозии порчаи BC: ".$BC;
echo "<br>Суммаи порчаҳои AC ва BC: ".$jamm;
?>
*/