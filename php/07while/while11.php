<?php
$sum = $K = 0;
$N = mt_rand(1, 100);
echo "N = ".$N;
while($sum < $N){
	$K++;
	$sum += $K;
}
echo "<br>K = ".$K;
echo "<br>Cумма = ".$sum;
?>