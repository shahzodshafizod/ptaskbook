<?php
$hisob = 1;
$K = mt_rand(1, 20);
echo "K = ".$K."<br>";
$N = mt_rand(2, 30);
echo "Миқдори ададҳои маҷмӯъ = ".$N."<br>";
while($hisob <= $N){
	$A = mt_rand(1, 100);
	echo "A".$hisob." = ".$A."<br>";
	$dar = 1;
	for($i = 1; $i <= $K; $i++)
		$dar *= $A;
	echo "".$A."^".$K." = ".$dar."<br>";
	$hisob++;
}
?>