<?php
$x1 = mt_rand(-100, 100) / 4;
$x2 = mt_rand(-100, 100) / 4;
$S = abs($x2 - $x1);
$x1 = number_format($x1, 2);
$x2 = number_format($x2, 2);
$S = number_format($S, 2);
echo "Первая точка: ".$x1;
echo "<br>Вторая точка: ".$x2;
echo "<br>Расстояние между двумя точками: ".$S;
?>

/*
<?php
$x1 = mt_rand(-100, 100) / 4;
$x2 = mt_rand(-100, 100) / 4;
$S = abs($x2 - $x1);
$x1 = number_format($x1, 2);
$x2 = number_format($x2, 2);
$S = number_format($S, 2);
echo "Нуқтаи якӯм: ".$x1;
echo "<br>Нуқтаи дуюм: ".$x2;
echo "<br>Масофаи байни ду нуқта: ".$S;
?>
*/