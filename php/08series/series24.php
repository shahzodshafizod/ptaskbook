<?php
$sum1 = 0;
$sum = 0;
$hisob = 2;
$mant = false;
$N = mt_rand(3, 30);
echo "Миқдори ададҳои маҷмӯъ = ".$N."<br>";
$A = mt_rand(1, 100);
echo "n1 = ".$A."<br>";
while($hisob <= $N){
	if($A == 0)
		$mant = true;
	$B = mt_rand(1, 100);
	echo "n".$hisob." = ".$B."<br>";
	if(($B != 0) && $mant)
		$sum += $B;
	else{
		$sum1 = $sum;
		$A = $B;
		$sum = 0;
	}
	$hisob++;
}
echo "Суммаи ададҳои байни нулҳо = ".$sum1;
?>