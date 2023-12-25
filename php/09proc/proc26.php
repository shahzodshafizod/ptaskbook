<?php
function IsPower5($K){
        $dar = 1;
        while($dar < $K)
                $dar *= 5;
        $nat = ($dar == $K);
        return $nat;
}
$hisob = 1;
$miq = 0;
while($hisob <= 10){
	$n = mt_rand(1, 1000000);
	echo "n".$hisob." = ".$n."<br>";
	if(IsPower5($n))
		$miq++;
	$hisob++;
}
echo "Миқдори дараҷаҳои адади панҷ: ".$miq;
?>