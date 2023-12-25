<?php
$a = mt_rand(1, 50);
$b = mt_rand(1, 50);
$c = sqrt($a*$a + $b*$b);
$P = $a + $b + $c;
$c = number_format($c, 2);
$P = number_format($P, 2);
echo "Первый катет треугольника: ".$a;
echo "<br>Второй катет треугольника: ".$b;
echo "<br>Гипотенуза треугольника: ".$c;
echo "<br>Периметр треугольника: ".$P;
?>

/*
<?php
$a = mt_rand(1, 50);
$b = mt_rand(1, 50);
$c = sqrt($a*$a + $b*$b);
$P = $a + $b + $c;
$c = number_format($c, 2);
$P = number_format($P, 2);
echo "Катети якӯми секунҷа: ".$a;
echo "<br>Катети дуюми секунҷа: ".$b;
echo "<br>Гипотенузаи секунҷа: ".$c;
echo "<br>Периметри секунҷа: ".$P;
?>
*/