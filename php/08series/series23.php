<?php
$hisob = 3;
$index = 0;
$mant = true;
$N = mt_rand(3, 30);
echo "Миқдори ададҳои маҷмӯъ = ".$N."<br>";
$A = mt_rand(1, 100);
echo "n1 = ".$A."<br>";
$B = mt_rand(1, 100);
echo "n2 = ".$B."<br>";
while($hisob <= $N){
	$C = mt_rand(1, 100);
	echo "n".$hisob." = ".$C."<br>";
	if((($B > $A) && ($B < $C) || ($B < $A) && ($B > $C)) && $mant){
		$index = $hisob;
		$mant = false;
	}
	$A = $B;
	$B = $C;
	$hisob++;
}
echo "index = ".$index;
?>