<?php
$zarb = 1;
$hisob = 1;
$N = mt_rand(1, 30);
echo "Миқдори элементҳои маҷмӯъ = ".$N."<br>";
while($hisob <= $N){
	do{
		$A = mt_rand(1, 100) / 4;
		$but = intval($A);
		$a = $A - $but;
	}
	while(($A <= 0) || ($a == 0));
	echo "A".$hisob." = ".$A."<br>";
	$zarb *= $a;
	echo "A".$hisob." = ".$a."<br><br>";
	$hisob++;
}
echo "Зарби қисми даҳии ададҳо = ".$zarb;
?>