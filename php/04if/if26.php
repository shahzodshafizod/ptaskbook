<?php
$x = mt_rand(-100, 100);
if($x <= 0)
	$y = -$x;
else if($x >= 2)
	$y = 4;
else
	$y = $x * $x;
echo "x = ".$x;
echo "<br>y = ".$y;
?>