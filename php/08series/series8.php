<?php
$K = 0;
$hisob = 1;
$N = mt_rand(1, 30);
echo "Миқдори элементҳои маҷмӯъ = ".$N."<br>";
while($hisob <= $N){
	$A = mt_rand(1, 100);
	echo "A".$hisob." = ".$A."<br>";
	if(($A % 2) == 0){
		$K++;
	}
	$hisob++;
}
echo "Миқдори ададҳои ҷуфт = ".$K;
?>