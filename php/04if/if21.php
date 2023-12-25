<?php
$x = mt_rand(-10, 10);
$y = mt_rand(-10, 10);
echo "x = ".$x;
echo "<br>y = ".$y."<br>Натиҷа: ";
if(($x == 0) && ($y == 0))
	echo "0";
else if($y == 0)
	echo "1";
else if($x == 0)
	echo "2";
else
	echo "3";
?>