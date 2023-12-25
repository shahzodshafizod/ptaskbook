<?php
$K = $sum = 0;
$A = mt_rand(2, 20);
echo "A = ".$A;
while($sum <= $A){
	$K++;
	$sum += 1 / $K;
}
echo "<br>K = ".$K;
echo "<br>Cумма = ".$sum;
?>