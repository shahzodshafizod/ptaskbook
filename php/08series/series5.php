<?php
$hisob = 1;
$sum = 0;
$N = mt_rand(1, 30);
echo "Миқдори элементҳои маҷмӯъ = ".$N."<br>";
while($hisob <= $N){
	do{
		$A = mt_rand(-100, 100) / 4;
	}
	while($A == 0);
	echo "A".$hisob." = ".$A."<br>";
	$but = intval($A); //Қисми бутуни адад
	echo "A".$hisob." = ".$but."<br>";
	$sum += $but;
	$hisob++;
}
echo "Ҷамъи қисми бутуни ададҳо = ".$sum;
?>