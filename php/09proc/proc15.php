<?php
function ShiftLeft3(&$A, &$B, &$C){
        $D = $A;
        $A = $B;
        $B = $C;
        $C = $D;
}
$hisob = 1;
while($hisob <= 2){
	$A = mt_rand(-100, 100);
	$B = mt_rand(-100, 100);
	$C = mt_rand(-100, 100);
	echo "A".$hisob." = ".$A;
	echo "<br>B".$hisob." = ".$B;
	echo "<br>C".$hisob." = ".$C;
	ShiftLeft3($A, $B, $C);
	echo "<br>Ададҳо пас аз иҷрои амалиёт:";
	echo "<br>A".$hisob." = ".$A;
	echo "<br>B".$hisob." = ".$B;
	echo "<br>C".$hisob." = ".$C."<br><br>";
	$hisob++;
}
?>