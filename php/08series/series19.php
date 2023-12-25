<?php
$hisob = 2;
$miq = 0;
$N = mt_rand(1, 30);
echo "Миқдори ададҳои маҷмӯъ = ".$N."<br>";
$A = mt_rand(1, 100);
echo "n1 = ".$A."<br>";
while($hisob <= $N){
	$B = mt_rand(1, 100);
	echo "n".$hisob." = ".$B."<br>";
	if($B > $A){
		$miq++;
	}
	$A = $B;
	$hisob++;
}
echo "Миқдори ададҳои талабкардашуда = ".$miq;
?>