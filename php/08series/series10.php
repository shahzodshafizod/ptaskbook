<?php
function boolval($nat){
	return ($nat? "true": "false");
}
$hisob = 1;
$mant = false;
$N = mt_rand(1, 30);
echo "Миқдори элементҳои маҷмӯъ = ".$N."<br>";
while($hisob <= $N){
	$A = mt_rand(-100, 100);
	echo "A".$hisob." = ".$A."<br>";
	if($A > 0)
		$mant = true;
	$hisob++;
}
echo "Hатиҷа: ".boolval($mant);
?>