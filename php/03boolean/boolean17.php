<?php
function boolval($nat){
	return ($nat? 'true': 'false');
}
$N = mt_rand(1, 10000);
$mant = (($N % 2) == 1) && ($N > 100) && ($N < 1000);
echo "Данное число: ".$N;
echo "<br>Нечетность и трехзначность данного число: ".boolval($mant);
?>