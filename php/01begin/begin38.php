<?php
do{
	$A = mt_rand(-100, 100);
}
while($A == 0);
$B = mt_rand(-100, 100);
$x = -$B / $A;
$x = number_format($x, 2);
echo "Первый коэффициент: ".$A;
echo "<br>Второй коэффициент: ".$B;
echo "<br>Корень уравнения: ".$x;
?>

/*
<?php
do{
	$A = mt_rand(-100, 100);
}
while($A == 0);
$B = mt_rand(-100, 100);
$x = -$B / $A;
$x = number_format($x, 2);
echo "Коэффитсиенти якӯм: ".$A;
echo "<br>Коэффитсиенти дуюм: ".$B;
echo "<br>Решаи муодила: ".$x;
?>
*/