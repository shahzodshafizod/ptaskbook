<?php
$A1 = 1;
$A2 = 2;
$K = 3;
$esp = array(0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9);
$Ak = ($A1 + 2 * $A2) / 3;
$index = mt_rand(0, 8);
while(abs($Ak - $A2) >= $esp[$index]){
	$A1 = $A2;
	$A2 = $Ak;
	$Ak = ($A1 + 2 * $A2) / 3;
	$K++;
}
echo "<br>K = ".$K;
echo "<br>\nA[K-1] = ".number_format($A2, 2);
echo "<br>\nA[K] = ".number_format($Ak, 2);
?>