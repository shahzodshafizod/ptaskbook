<?php
$miq = 0;
$sum = 0;
$i = 1;
$j = 1;
$K = mt_rand(1, 30);
echo "Миқдори маҷмӯъҳо = ".$K."<br>";
while($i <= $K){
	echo "Nabor".$i."<br>";
	do{
		$n = mt_rand(-100, 100);
		echo "n".$j." = ".$n."<br>";
		if($n != 0)
			$miq++;
		$j++;
	}
	while($n != 0);
	echo "Миқдори ададҳои маҷмӯъ = ".$miq."<br>";
	$sum += $miq;
	$miq = 0;
	$j = 1;
	$i++;
}
echo "Миқдори ададҳои маҷмӯъҳо = ".$sum;
 ?>