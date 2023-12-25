<?php
$hisob = 1;
$i = 3;
$raq = 0;
$miq = 2;
$mant = true;
$K = mt_rand(1, 30);
echo "Миқдори маҷмӯъҳо = ".$K."<br>";
while($hisob <= $K){
echo "Nabor".$hisob."<br>";
do{
	$a = mt_rand(-100, 100);
	echo "n1 = ".$a."<br>";
	$b = mt_rand(-100, 100);
	echo "n2 = ".$b."<br>";
}
while(($a == 0) || ($b == 0));
do{
	$c = mt_rand(-100, 100);
	echo "n".$i." = ".$c."<br>";
	if($c != 0){
		if((($b>$a) && ($b<$c) || ($b<$a) && ($b>$c)) && $mant){
			$raq = $i;
			$mant = false;
		}
		else
			$miq++;
	}
	$a = $b;
	$b = $c;
	$i++;
}
while($c != 0);
	if($mant)
		echo "Миқдори ададҳои маҷмӯъ = ".$miq;
	else
		echo $raq;
	$hisob++;
	$mant = true;
	$raq = 0;
	$miq = 2;
	$i = 3;
}
?>