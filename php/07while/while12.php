<?php
$K = $sum = 0;
$N = mt_rand(1, 100);
echo "N = ".$N;
while($sum <= $N){
	$K++;
	$sum += $K;
}
echo "<br>K = ".($K - 1);
echo "<br>Cумма = ".($sum - $K);
?>