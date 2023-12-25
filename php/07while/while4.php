<?php
function boolval($nat){
	return ($nat? "true": "false");
}
$dar = 1;
$N = mt_rand(1, 100);
echo "N = ".$N;
while($dar < $N)
	$dar *= 3;
echo "<br>Натиҷа: ".boolval($dar == $N);
?>