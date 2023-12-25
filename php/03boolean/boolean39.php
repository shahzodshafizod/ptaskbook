<?php
function boolval($nat){
	return ($nat? "true": "false");
}
$x1 = mt_rand(1, 8);
$y1 = mt_rand(1, 8);
$x2 = mt_rand(1, 8);
$y2 = mt_rand(1, 8);
echo "Координаты первого поля шахматной доски:";
echo "<br>x1 = ".$x1."<br>y1 = ".$y1;
echo "<br>Координаты второго поля шахматной доски:";
echo "<br>x2 = ".$x2."<br>y2 = ".$y2;
$mant = ($x1==$x2) || ($y1==$y2) || (abs($x1-$x2) == abs($y1-$y2));
echo "<br>Движение Ферзя за один ход с одного поля на другое:  ".boolval($mant);
?>