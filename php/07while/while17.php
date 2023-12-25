<?php
$N = mt_rand(1, 1000000);
echo "N = ".$N."<br>";
while($N > 0){
	$baq = $N % 10;
	$N = intval($N / 10);
	echo $baq."<br>";
}
?>