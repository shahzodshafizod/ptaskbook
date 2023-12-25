<?php
$kv = 0;
$N = mt_rand(1, 50);
echo "Ҳудуди давр: ".$N."<br>";
for($i = 1; $i <= 2*$N-1; $i += 2){
	$kv += $i;
	echo $i." = ".$kv."<br>";
}
function boolval($nat){
	return ($nat? "true": "false");
}
echo "Дурустии ".$kv." = ".$N."^2:  ".boolval($kv == $N*$N);
?>