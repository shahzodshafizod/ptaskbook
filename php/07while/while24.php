﻿<?php
function boolval($nat){
	return ($nat? 'true': 'false');
}
$F1 = $F2 = 1;
$N = mt_rand(2, 100);
echo "N = ".$N;
$Fk = $F1 + $F2;
$F1 = $F2;
$F2 = $Fk;
while($N > $Fk){
	$Fk = $F1 + $F2;
	$F1 = $F2;
	$F2 = $Fk;
}
echo "<br>Натиҷа: ".boolval($N == $Fk);
?>