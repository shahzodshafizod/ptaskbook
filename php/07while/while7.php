<?php
$K = 1;
$N = mt_rand(1, 100);
echo "N = ".$N;
while($K * $K <= $N)
	$K++;
echo "<br>K = ".$K;
?>