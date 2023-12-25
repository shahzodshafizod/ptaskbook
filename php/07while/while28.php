<?php
$A1 = 2;
$K = 2;
$esp = array(0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9);
$index = mt_rand(0, 8);
$Ak = 2 + 1 / $A1;
while(abs($Ak - $A1) >= $esp[$index]){
	$A1 = $Ak;
	$Ak = 2 + 1 / $A1;
	$K++;
}
echo "<br>K = ".$K;
echo "<br>\nA[K-1] = ".number_format($A1, 2);
echo "<br>\nA[K+1] = ".number_format($Ak, 2);
?>