<?php
$A = mt_rand(1, 100);
$B = mt_rand(1, 100);
echo "A = ".$A;
echo "<br>B = ".$B;
echo "<br>КТУ(".$A.", ".$B.") = ";
while(($A != 0) && ($B != 0))
	if($A >= $B)
		$A %= $B;
	else
		$B %= $A;
echo ($A + $B);
?>