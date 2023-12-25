<?php
function boolval($nat){
	return ($nat? 'true': 'false');
}
$A = mt_rand(-100, 100);
$B = mt_rand(-100, 100);
$C = mt_rand(-100, 100);
$mant = ($A>0)&&($B<0)&&($C<0) || ($A<0)&&($B>0)&&($C<0) || ($A<0)&&($B<0)&&($C>0);
echo "Первое число: ".$A;
echo "<br>Второе число: ".$B;
echo "<br>Третье число: ".$C;
echo "<br>Положительность только одного числа: ".boolval($mant);
?>