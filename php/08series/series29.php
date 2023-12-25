<?php
$i = 1;
$j = 1;
$sum = 0;
$K = mt_rand(1, 30);
echo "Миқдори маҷмӯъҳо = ".$K."<br>";
$N = mt_rand(1, 30);
echo "Миқдори ададҳои маҷмӯъҳо = ".$N."<br>";
while($i <= $K){
	echo "Nabor".$i."<br>";
	while($j <= $N){
		$n = mt_rand(-100, 100);
		echo "n".$j." = ".$n."<br>";
		$sum += $n;
		$j++;
	}
	$j = 1;
	$i++;
}
echo "Ҷамъи ҳамаи ададҳо = ".$sum;
?>