<?php
function boolval($nat){
	return ($nat? 'true': 'false');
}
$N = mt_rand(1000, 9999);
$a = intval($N / 1000);
$b = intval($N % 1000 / 100);
$c = intval($N % 100 / 10);
$d = $N % 10;
$mant = ($a == $d) && ($c == $b);
echo "Четырехзначное число: ".$N;
echo "<br>Разяд тысяч: ".$a;
echo "<br>Разряд сотен: ".$b;
echo "<br>Разяд десятков: ".$c;
echo "<br>Разряд единиц: ".$d;
echo "<br>Единочитаемость числа с обоих сторон:  ".boolval($mant);
?>