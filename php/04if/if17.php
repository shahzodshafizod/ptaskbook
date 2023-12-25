<?php
$A = mt_rand(-100, 100);
$B = mt_rand(-100, 100);
$C = mt_rand(-100, 100);
echo "Aдади якӯм: ".$A;
echo "<br>Aдади дуюм: ".$B;
echo "<br>Aдади сеюм: ".$C;
if(($A < $B) && ($B < $C) || ($A > $B) && ($B > $C)){
	$A *= 2;
	$B *= 2;
	$C *= 2;
}
else{
	$A = -$A;
	$B = -$B;
	$C = -$C;
}
echo "<br>Aдади якӯм: ".$A;
echo "<br>Aдади дуюм: ".$B;
echo "<br>Aдади сеюм: ".$C;
?>