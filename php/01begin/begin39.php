﻿<?php
do{
	$A = mt_rand(-100, 100);
}
while($A == 0);
$B = mt_rand(-100, 100);
$C = mt_rand(-100, 100);
$D = $B * $B - 4 * $A * $C;
$x1 = (-$B + sqrt($D)) / (2 * $A);
$x2 = (-$B - sqrt($D)) / (2 * $A);
echo "Первый коэффициент: ".$A;
echo "<br>Второй коэффициент: ".$B;
echo "<br>Третий коэффитсиент: ".$C;
echo "<br>Дискриминант: ".$D;
echo "<br>Первая корень: ".$x1;
echo "<br>Вторая корень: ".$x2;
?>

/*
<?php
do{
	$A = mt_rand(-100, 100);
}
while($A == 0);
$B = mt_rand(-100, 100);
$C = mt_rand(-100, 100);
$D = $B * $B - 4 * $A * $C;
$x1 = (-$B + sqrt($D)) / (2 * $A);
$x2 = (-$B - sqrt($D)) / (2 * $A);
echo "Коэффитсиенти якӯм: ".$A;
echo "<br>Коэффитсиенти дуюм: ".$B;
echo "<br>Коэффитсиенти сеюм: ".$C;
echo "<br>Дискриминант: ".$D;
echo "<br>Решаи якӯм: ".$x1;
echo "<br>Решаи дуюм: ".$x2;
?>
*/