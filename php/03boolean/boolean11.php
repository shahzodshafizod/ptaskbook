<?php
function boolval($nat){
	return ($nat? 'true': 'false');
}
$A = mt_rand(-100, 100);
$B = mt_rand(-100, 100);
$mant = ((abs($A) % 2) == (abs($B) % 2));
echo "Первое число: ".$A;
echo "<br>Второе число: ".$B;
echo "<br>Единочетность обоих чисел: ".boolval($mant);
?>