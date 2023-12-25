<?php
$miq = 0;
$hisob = 0;
$K = mt_rand(1, 100);
echo "K = ".$K."<br>";
do{
	$N = mt_rand(-100, 100);
	if($N != 0){
		echo "N".($hisob + 1)." = ".$N."<br>";
		if(($N < $K) && ($N != 0))
			$miq++;
		$hisob++;
	}
}
while($N != 0);
echo "Миқдори ададҳои аз K хурд = ".$miq;
?>