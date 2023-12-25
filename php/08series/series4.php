<?php
$hisob = 1;
$sum = 0;
$zarb = 1;
$N = mt_rand(1, 30);
echo "Миқдори элементҳои маҷмӯъ = ".$N."<br>";
while($hisob <= $N){
	$A = mt_rand(1, 100) / 4;
	$A = number_format($A, 2);
	echo "A".$hisob." = ".$A."<br>";
	$sum += $A;
	$zarb *= $A;
	$hisob++;
}
echo "Ҷамъи ададҳо = ".$sum;
echo "<br>Зарби ададҳо = ".$zarb;
?>