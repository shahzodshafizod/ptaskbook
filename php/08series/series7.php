<?php
$hisob = 1;
$sum = 0;
$N = mt_rand(1, 30);
echo "Миқдори элементҳои маҷмӯъ = ".$N."<br>";
while($hisob <= $N){
	$A = mt_rand(1, 100) / 4;
	echo "A".$hisob." = ".$A."<br>";
	$nat = ceil($A);
	echo "A".$hisob." = ".$nat."<br><br>";
	$sum += $nat;
	$hisob++;
}
echo "Ҷамъи ададҳои яклухткардашуда = ".$sum;
?>