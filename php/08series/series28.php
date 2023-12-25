<?php
$hisob = 1;
$N = mt_rand(2, 30);
echo "Миқдори ададҳои маҷмӯъ = ".$N."<br>";
while($hisob <= $N){
	$A = mt_rand(1, 100);
	echo "A".$hisob." = ".$A."<br>";
	$dar = 1;
	for($i = 1; $i <= ($N - $hisob + 1); $i++)
		$dar *= $A;
	echo "".$A."^".($N - $hisob + 1)." = ".$dar."<br>";
	$hisob++;
}
?>