<?php
function boolval($nat){
	return ($nat? 'true': 'false');
}
$i = 1;
$j = 1;
$sum = 0;
$mant = false;
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
		if($n == 2)
			$mant = true;
		$j++;
	}
	if($mant)
		echo "Сумма = ".$sum."<br>";
	else
		echo "Натиҷа: ".boolval($mant)."<br><br>";
	$sum = 0;
	$mant = false;
	$j = 1;
	$i++;
}
 ?>