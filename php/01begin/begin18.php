<?php
$A = mt_rand(-100, 100) / 4;
$B = mt_rand(-100, 100) / 4;
$C = mt_rand(-100, 100) / 4;
$AC = abs($C - $A);
$BC = abs($C - $B);
$zarb = $AC * $BC;
$A = number_format($A, 2);
$B = number_format($B, 2);
$C = number_format($C, 2);
$AC = number_format($AC, 2);
$BC = number_format($BC, 2);
$zarb = number_format($zarb, 2);
echo "Первая точка: ".$A;
echo "<br>Вторая точка: ".$B;
echo "<br>Третья точка: ".$C;
echo "<br>Длина отрезка AC: ".$AC;
echo "<br>Длина отрезка BC: ".$BC;
echo "<br>Проиведение длин отрезков AC и BC: ".$zarb;
?>

/*
<?php
$A = mt_rand(-100, 100) / 4;
$B = mt_rand(-100, 100) / 4;
$C = mt_rand(-100, 100) / 4;
$AC = abs($C - $A);
$BC = abs($C - $B);
$zarb = $AC * $BC;
$A = number_format($A, 2);
$B = number_format($B, 2);
$C = number_format($C, 2);
$AC = number_format($AC, 2);
$BC = number_format($BC, 2);
$zarb = number_format($zarb, 2);
echo "Нуқтаи якӯм: ".$A;
echo "<br>Нуқтаи дуюм: ".$B;
echo "<br>Нуқтаи сеюм: ".$C;
echo "<br>Дарозии порчаи AC: ".$AC;
echo "<br>Дарозии порчаи BC: ".$BC;
echo "<br>ҲосилиЗарби порчаҳои AC ва BC: ".$zarb;
?>
*/