<?php
$x1 = mt_rand(-100, 100) / 4;
$x2 = mt_rand(-100, 100) / 4;
$x3 = mt_rand(-100, 100) / 4;
$y1 = mt_rand(-100, 100) / 4;
$y2 = mt_rand(-100, 100) / 4;
$y3 = mt_rand(-100, 100) / 4;
$a = sqrt(pow($x2 - $x1, 2) + pow($y2 - $y1, 2));
$b = sqrt(pow($x3 - $x2, 2) + pow($y3 - $y2, 2));
$c = sqrt(pow($x3 - $x1, 2) + pow($y3 - $y1, 2));
$P = $a + $b + $c;
$p = $P / 2;
$S = sqrt($p * ($p - $a) * ($p - $b) * ($p - $c));
$x1 = number_format($x1, 2);
$x2 = number_format($x2, 2);
$x3 = number_format($x3, 2);
$y1 = number_format($y1, 2);
$y2 = number_format($y2, 2);
$y3 = number_format($y3, 2);
$a = number_format($a, 2);
$b = number_format($b, 2);
$c = number_format($c, 2);
$P = number_format($P, 2);
$p = number_format($p, 2);
$S = number_format($S, 2);
echo "Координаты первой точки:";
echo "<br>x1 = ".$x1."<br>y1 = ".$y1;
echo "<br>Координаты второй точки:";
echo "<br>x2 = ".$x2."<br>y2 = ".$y2;
echo "<br>Координаты третьей точки:";
echo "<br>x3 = ".$x3."<br>y3 = ".$y3;
echo "<br>Первая сторона треугольника: ".$a;
echo "<br>Вторая сторона треугольника: ".$b;
echo "<br>Третья сторона треугольника: ".$c;
echo "<br>Периметр: ".$P;
echo "<br>ПолуПериметр: ".$p;
echo "<br>Площадь треугольника по формуле Герона: ".$S;
?>

/*
<?php
$x1 = mt_rand(-100, 100) / 4;
$x2 = mt_rand(-100, 100) / 4;
$x3 = mt_rand(-100, 100) / 4;
$y1 = mt_rand(-100, 100) / 4;
$y2 = mt_rand(-100, 100) / 4;
$y3 = mt_rand(-100, 100) / 4;
$a = sqrt(pow($x2 - $x1, 2) + pow($y2 - $y1, 2));
$b = sqrt(pow($x3 - $x2, 2) + pow($y3 - $y2, 2));
$c = sqrt(pow($x3 - $x1, 2) + pow($y3 - $y1, 2));
$P = $a + $b + $c;
$p = $P / 2;
$S = sqrt($p * ($p - $a) * ($p - $b) * ($p - $c));
$x1 = number_format($x1, 2);
$x2 = number_format($x2, 2);
$x3 = number_format($x3, 2);
$y1 = number_format($y1, 2);
$y2 = number_format($y2, 2);
$y3 = number_format($y3, 2);
$a = number_format($a, 2);
$b = number_format($b, 2);
$c = number_format($c, 2);
$P = number_format($P, 2);
$p = number_format($p, 2);
$S = number_format($S, 2);
echo "Координатаҳои нуқтаи якӯм:";
echo "<br>x1 = ".$x1."<br>y1 = ".$y1;
echo "<br>Координатаҳои нуқтаи дуюм:";
echo "<br>x2 = ".$x2."<br>y2 = ".$y2;
echo "<br>Координатаҳи нуқтаи сеюм:";
echo "<br>x3 = ".$x3."<br>y3 = ".$y3;
echo "<br>Тарафи якӯми секунҷа: ".$a;
echo "<br>Тарафи дуюми секунҷа: ".$b;
echo "<br>Тарафи сеюми секунҷа: ".$c;
echo "<br>Периметр: ".$P;
echo "<br>НимПериметр: ".$p;
echo "<br>Масоҳати секунҷа аз рӯи формулаи Ҳерон: ".$S;
?>
*/