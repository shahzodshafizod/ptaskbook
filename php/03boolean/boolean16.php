<?php
function boolval($nat){
	return ($nat? 'true': 'false');
}
$N = mt_rand(1, 1000);
$mant = (($N % 2) == 0) && ($N > 9) && ($N < 99);
echo "Данное число: ".$N;
echo "<br>Четность и двухзначность данного числа: ".boolval($mant);
?>