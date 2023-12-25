<?php
$x = mt_rand(-100, 100) / 4;
echo "x = ".number_format($x);
$x = intval($x);
if ($x < 0)
	$y = 0;
else if(($x % 2) == 0)
	$y = 1;
else
	$y = -1;
echo "<br>y = ".$y;
?>