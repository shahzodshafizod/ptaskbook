<?php
const PI = 3.14;
function RingS($R1, $R2){
        $S1 = PI * $R1 * $R1;
        $S2 = PI * $R2 * $R2;
        $nat = $S1 - $S2;
		$nat = number_format($nat, 2);
        return $nat;
}
$hisob = 1;
while($hisob <= 3){
	do{
		$r1 = mt_rand(1, 100) / 4;
		$r2 = mt_rand(1, 100) / 4;
	}
	while(($r1 <= $r2) || ($r1 == 0) || ($r2 == 0));	
	$natija = RingS($r1, $r2);
	$r1 = number_format($r1, 2);
	$r2 = number_format($r2, 2);
	echo "Paдиуси давраи якӯм: ".$r1;
	echo "<br>Pадиуси давраи дуюм: ".$r2;
	echo "<br>Масоҳати ҳалқа: ".$natija."<br><br>";
	$hisob++;
}
?>