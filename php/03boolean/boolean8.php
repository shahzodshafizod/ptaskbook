<?php
function boolval($nat){ 
	return ($nat? 'true': 'false');
}
$A = mt_rand(-100, 100);
$B = mt_rand(-100, 100);
$mant = (($A % 2) == (1 || -1) || ($B % 2) == (1 || -1));
echo "Первое число: ".$A;
echo "<br>Второе число: ".$B;
echo "<br>Нечетность обоих чисел: ".boolval($mant);
?>