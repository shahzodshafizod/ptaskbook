<?php
$dar = 1;
$K = 0;
$N = mt_rand(1, 100);
echo "N = ".$N;
while($dar < $N){
	$dar *= 3;
	$K++;
}
echo "<br>K = ".(--$K);
?>