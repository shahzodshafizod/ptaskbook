<?php
function boolval($nat){
	return ($nat? "true": "false");
}
$mant = false;
$N = mt_rand(1, 1000);
echo "N = ".$N;
while($N > 0){
	$baq = $N % 10;
	if(($baq % 2) == 1)
		$mant = true;
	$N = intval($N / 10);
}
echo "<br>Натиҷа: ".boolval($mant);
?>