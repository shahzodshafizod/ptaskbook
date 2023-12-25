<?php
$hisob = 2;
$index = 0;
$mant = true;
$N = mt_rand(1, 30);
echo "Миқдори ададҳои маҷмӯъ = ".$N."<br>";
$A = mt_rand(1, 100);
echo "n1 = ".$A."<br>";
while($hisob <= $N){
	$B = mt_rand(1, 100);
	echo "n".$hisob." = ".$B."<br>";
	if(($A < $B) && $mant){
		$index = $hisob;
		$mant = false;
	}
	$A = $B;
	$hisob++;
}
echo "index = ".$index;
?>