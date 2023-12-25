<?php
function boolval($nat){ 
	return ($nat? "true": "false");
}
$A = mt_rand(-100, 100);
$B = mt_rand(-100, 100);
$mant = (($A > 2) && ($B <= 3));
echo "Первое число: ".$A;
echo "<br>Второе число: ".$B;
echo "<br>Справедливость неравенства A>2 и B<=3 :  ".boolval($mant);
?>