<?php
function boolval($nat){
	return ($nat? 'true': 'false');
}
$N = mt_rand(100, 999);
$a = intval($N / 100);
$b = intval($N / 10) % 10;
$c = $N % 10;
$mant = ($a < $b) && ($b < $c);
echo "Трехзначное число: ".$N;
echo "<br>Разряд сотен: ".$a;
echo "<br>Разряд десятков: ".$b;
echo "<br>Разряд единиц: ".$c;
echo "<br>Возрастение цифр данного числа:  ".boolval($mant);
?>