<?php
$N = mt_rand(1, 30);
echo "Миқдори ададҳои маҷмӯъ = ".$N."<br>";
$A = mt_rand(-100, 100);
$k = 2;
$hisob = 2;
echo "n1 = ".$A."<br>";
while($hisob <= $N){
	$B = mt_rand(-100, 100);
	if($A != $B){
		echo "n".$k." = ".$B."<br>";
		$k++;
	}
	else
		echo "index = ".$hisob."<br>";
	$hisob++;
}
?>