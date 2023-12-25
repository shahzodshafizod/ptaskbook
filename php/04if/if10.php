<?php
$a = mt_rand(-100, 100);
$b = mt_rand(-100, 100);
echo "A = ".$a;
echo "<br>B = ".$b;
$sum = $a + $b;
if($a != $b)
	$a = $b = $sum;
else
	$a = $b = 0;
echo "<br>Ададҳо пас аз иҷрои шарт:";
echo "<br>A = ".$a;
echo "<br>B = ".$b;
?>