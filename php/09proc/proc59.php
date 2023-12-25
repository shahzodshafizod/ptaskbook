<?php
function Leng($xA, $yA, $xB, $yB){
        return number_format(sqrt(pow($xA - $xB, 2) + pow($yA - $yB, 2)), 2);
}
function Area($xA, $yA, $xB, $yB, $xC, $yC){
        $a = Leng($xA, $yA, $xB, $yB);
        $b = Leng($xB, $yB, $xC, $yC);
        $c = Leng($xA, $yA, $xC, $yC);
        $P = $a + $b + $c;
        $p = $P / 2;
        $S = sqrt($p * ($p - $a) * ($p - $b) * ($p - $c));
		$S = number_format($S, 2);
        return $S;
}
function Dist($xP, $yP, $xA, $yA, $xB, $yB){
        return number_format(2 * Area($xP, $yP, $xA, $yA, $xB, $yB) / Leng($xA, $yA, $xB, $yB), 2);
}
$xP = mt_rand(1, 100);
$yP = mt_rand(1, 100);
$xA = mt_rand(1, 100);
$yA = mt_rand(1, 100);
$xB = mt_rand(1, 100);
$yB = mt_rand(1, 100);
$xC = mt_rand(1, 100);
$yC = mt_rand(1, 100);
echo "xP = ".$xP;
echo "<br>yP = ".$yP;
echo "<br>xA = ".$xA;
echo "<br>yA = ".$yA;
echo "<br>xB = ".$xB;
echo "<br>yB = ".$yB;
echo "<br>xC = ".$xC;
echo "<br>yC = ".$yC;
$PAB = Dist($xP, $yP, $xA, $yA, $xB, $yB);
$PAC = Dist($xP, $yP, $xA, $yA, $xC, $yC);
$PBC = Dist($xP, $yP, $xB, $yB, $xC, $yC);
echo "<br>Dist(P, AB) = ".$PAB;
echo "<br>Dist(P, AC) = ".$PAC;
echo "<br>Dist(P, BC) = ".$PBC;
?>