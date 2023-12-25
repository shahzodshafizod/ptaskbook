<?php
$raq = 0;
$hisob = 1;
$mant = true;
$K = mt_rand(1, 100);
echo "K = ".$K."<br>";
do{
	$N = mt_rand(-100, 100);
	if($N != 0){
		echo "N".$hisob." = ".$N."<br>";
		if(($N > $K) && ($mant) && ($N != 0)){
			$raq = $hisob;
			$mant = false;
		}
		$hisob++;
	}
}
while($N != 0);
echo "Рақами адади аввалини аз К калон = ".$raq;
?>