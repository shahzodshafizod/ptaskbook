<?php
$x = mt_rand(-100, 100);
if(($x < -2) && ($x > 2))
	$y = 2 * $x;
else
	$y = -3 * $x;
echo "x = ".$x;
echo "<br>y = ".$y;
?>