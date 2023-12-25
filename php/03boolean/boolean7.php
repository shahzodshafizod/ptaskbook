<?php
function boolval($nat){ 
	return ($nat? 'true': 'false');
}
$A = mt_rand(-100, 100);
$B = mt_rand(-100, 100);
$C = mt_rand(-100, 100);
$mant = (($A<$B) && ($B<$C) || ($A>$B) && ($B>$C));
echo "Первое число: ".$A;
echo "<br>Второе число: ".$B;
echo "<br>Третье число: ".$C;
echo "<br>Среди чисел A и C раполжение числа C:  ".boolval($mant);
?>