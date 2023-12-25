<?php
function boolval($nat){
	return ($nat? "true": "false");
}
$i = 2;
$mant = true;
$N = mt_rand(2, 100);
echo "N = ".$N;
while($i < sqrt($N)){
	if(($N % $i) == 0)
		$mant = false;
	$i++;
}
echo "<br>Дурустии оддӣ будани адад: ".boolval($mant);
?>