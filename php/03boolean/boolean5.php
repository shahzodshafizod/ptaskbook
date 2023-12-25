<?php
function boolval($nat){ 
	return ($nat? "true": "false");
}
$A = mt_rand(-100, 100);
$B = mt_rand(-100, 100);
$mant = (($A >= 0) || ($B < -2));
echo "Первое число: ".$A;
echo "<br>Второе число: ".$B;
echo "<br>Справедливость неравенства A>=0 или B<-2 :  ".boolval($mant);
?>