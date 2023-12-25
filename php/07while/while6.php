<?php
$dufact = 1;
$N = mt_rand(1, 20);
echo "ҳудуди факториали дукарата: ".$N;
while($N >= 1){
	$dufact *= $N;
	$N -= 2;
}
echo "<br>факториали дукарата = ".$dufact;
?>