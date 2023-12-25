<?php
$raq = 0;
$hisob = 1;
$K = mt_rand(1, 100);
echo "K = ".$K."<br>";
do{
	$N = mt_rand(-100, 100);
	if($N != 0){
		echo "N".$hisob." = ".$N."<br>";
		if(($N != 0) && ($N > $K))
			$raq = $hisob;
		$hisob++;
	}
}
while($N != 0);
echo "Рақами адади охирини аз К калон = ".$raq;
?>