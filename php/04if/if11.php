<?php
$A = mt_rand(-100, 100);
$B = mt_rand(-100, 100);
echo "A = ".$A;
echo "<br>B = ".$B;
if($A < $B)
	$A = $B;
else if($A > $B)
	$B = $A;
else
	$A = $B = 0;
echo "<br>Қиматҳои нави тағйирёбандаҳо:";
echo "<br>A = ".$A;
echo "<br>B = ".$B;
?>