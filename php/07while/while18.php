<?php
$sum = $miq = 0;
$N = mt_rand(1, 1000000);
echo "N = ".$N;
while($N > 0){
	$baq = $N % 10;
	$sum += $baq;
	$miq++;
	$N = intval($N / 10);
}
echo "<br>Миқдори рақамҳо = ".$miq;
echo "<br>Суммаи рақамҳо = ".$sum;
?>