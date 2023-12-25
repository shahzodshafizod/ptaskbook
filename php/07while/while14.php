<?php
$K = $sum = 1;
$A = mt_rand(1, 20);
echo "A = ".$A;
while($sum < $A){
	$K++;
	$sum += 1 / $K;
}
echo "<br>K = ".($K - 1);
echo "<br>Cумма = ".($sum -= 1 / $K);
?>