<?php
$x1 = mt_rand(-100, 100) / 4;
$x2 = mt_rand(-100, 100) / 4;
$y1 = mt_rand(-100, 100) / 4;
$y2 = mt_rand(-100, 100) / 4;
$a = abs($x2 - $x1);
$b = abs($y2 - $y1);
$P = 2 * ($a + $b);
$S = $a * $b;
$x1 = number_format($x1, 2);
$x2 = number_format($x2, 2);
$y1 = number_format($y1, 2);
$y2 = number_format($y2, 2);
$a = number_format($a, 2);
$b = number_format($b, 2);
$P = number_format($P, 2);
$S = number_format($S, 2);
echo "Координаты первой точки:";
echo "<br>x1 = ".$x1."<br>y1 = ".$y1;
echo "<br>Координаты второй точки";
echo "<br>x2 = ".$x2."<br>y2 = ".$y2;
echo "<br>Первая сторона прямоугольника: ".$a;
echo "<br>Вторая сторона прямоугольника: ".$b;
echo "<br>Периметр прямоугольника: ".$P;
echo "<br>Площадь прямоугольника: ".$S;
?>

/*
<?php
$x1 = mt_rand(-100, 100) / 4;
$x2 = mt_rand(-100, 100) / 4;
$y1 = mt_rand(-100, 100) / 4;
$y2 = mt_rand(-100, 100) / 4;
$a = abs($x2 - $x1);
$b = abs($y2 - $y1);
$P = 2 * ($a + $b);
$S = $a * $b;
$x1 = number_format($x1, 2);
$x2 = number_format($x2, 2);
$y1 = number_format($y1, 2);
$y2 = number_format($y2, 2);
$a = number_format($a, 2);
$b = number_format($b, 2);
$P = number_format($P, 2);
$S = number_format($S, 2);
echo "Координатаҳои нуқтаи якӯм:";
echo "<br>x1 = ".$x1."<br>y1 = ".$y1;
echo "<br>Координатаҳои нуқтаи дуюм";
echo "<br>x2 = ".$x2."<br>y2 = ".$y2;
echo "<br>Тарафи якӯми росткунҷа: ".$a;
echo "<br>Тарафи дуюми росткунҷа: ".$b;
echo "<br>Периметри росткунҷа: ".$P;
echo "<br>Масоҳати росткунҷа: ".$S;
?>
*/