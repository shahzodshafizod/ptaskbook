<?php
$raq = 0;
$i = 1;
$j = 1;
$mant = true;
$K = mt_rand(1, 30);
echo "Миқдори маҷмӯъҳо = ".$K."<br>";
$N = mt_rand(1, 30);
echo "Миқдори ададҳои маҷмӯъҳо = ".$N."<br>";
while($i <= $K){
	echo "Nabor".$i."<br>";
	while($j <= $N){
		$n = mt_rand(-100, 100);
		echo "n".$j." = ".$n."<br>";
		if(($n == 2) && $mant){
			$raq = $j;
			$mant = false;
		}
		$j++;
	}
	echo "Pақами адади ба 2 баробар = ".$raq."<br>";
	$raq = 0;
	$j = 1;
	$i++;
	$mant = true;
}
 ?>