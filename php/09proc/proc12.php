<?php
function Swap(&$A, &$B){
        if($A > $B){
                $C = $A;
                $A = $B;
                $B = $C;
        }
}
function SortInc3(&$A, &$B, &$C){
        Swap($A, $B);
        Swap($B, $C);
        Swap($A, $B);
}
$hisob = 1;
while($hisob <= 2){
	$A = mt_rand(-100, 100);
	$B = mt_rand(-100, 100);
	$C = mt_rand(-100, 100);
	echo "A".$hisob." = ".$A;
	echo "<br>B".$hisob." = ".$B;
	echo "<br>C".$hisob." = ".$C;
	SortInc3($A, $B, $C);
	echo "<br>Пас аз иҷрои амалиёт:";
	echo "<br>A".$hisob." = ".$A;
	echo "<br>B".$hisob." = ".$B;
	echo "<br>C".$hisob." = ".$C."<br><br>";
	$hisob++;
}
?>