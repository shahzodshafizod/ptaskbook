<?php
function boolval($nat){
	return ($nat? "true": "false");
}
$A = mt_rand(-100, 100);
$mant = (($A % 2) == (1 || -1));
echo "Данное число: ".$A;
echo "<br>Нечетность числа: ".boolval($mant);
?>