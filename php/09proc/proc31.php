<?php
function IsPalindrom($K){
        $adad = $K;
		$nav = 0;
        while($adad > 0){
                $nav = $nav * 10 + $adad % 10;
                $adad = intval($adad / 10);
        }
        $nat = ($nav == $K);
        return $nat;
}
$hisob = 1;
$miq = 0;
while($hisob <= 10){
	$n = mt_rand(1, 1000000);
	echo "n".$hisob." = ".$n."<br>";
	if(IsPalindrom($n))
		$miq++;
	$hisob++;
}
echo "Миқдори палиндромҳо: ".$miq;
?>