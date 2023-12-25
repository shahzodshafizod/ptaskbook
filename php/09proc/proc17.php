<?php
function RootsCount($A, $B, $C){
        $D = $B * $B - 4 * $A * $C;
        if($D > 0)
                $nat = 2;
        else if($D < 0)
                $nat = 0;
        else
                $nat = 1;
        return $nat;
}
$hisob = 1;
while($hisob <= 3){
	do{
		$A = mt_rand(-50, 50);
	}
	while($A == 0);
	$B = mt_rand(-50, 50);
	$C = mt_rand(-50, 50);
	echo "A = ".$A;
	echo "<br>B = ".$B;
	echo "<br>C = ".$C;
	$natija = RootsCount($A, $B, $C);
	echo "<br>Миқдори решаҳои муодила: ".$natija."<br><br>";
	$hisob++;
}
?>