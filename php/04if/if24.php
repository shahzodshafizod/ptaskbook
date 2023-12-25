<?php
$x = mt_rand(-100, 100);
if($x > 0)
	$y = 2 * sin($x);
else
	$y = 6 - $x;
echo "x = ".$x;
echo "<br>y = ".$y;
?>