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
function Heights($xA, $yA, $xB, $yB, $xC, $yC, &$HA, &$HB, &$HC){
		$HA = Dist($xA, $yA, $xC, $yC, $xB, $yB);
        $HB = Dist($xB, $yB, $xC, $yC, $xA, $yA);
        $HC = Dist($xC, $yC, $xA, $yA, $xB, $yB);
}
$xA = mt_rand(1, 100);
$yA = mt_rand(1, 100);
$xB = mt_rand(1, 100);
$yB = mt_rand(1, 100);
$xC = mt_rand(1, 100);
$yC = mt_rand(1, 100);
$xD = mt_rand(1, 100);
$yD = mt_rand(1, 100);
echo "<br>xA = ".$xA;
echo "<br>yA = ".$yA;
echo "<br>xB = ".$xB;
echo "<br>yB = ".$yB;
echo "<br>xC = ".$xC;
echo "<br>yC = ".$yC;
echo "<br>xD = ".$xD;
echo "<br>yD = ".$yD;
Heights($xA, $yA, $xB, $yB, $xC, $yC, $HA, $HB, $HC);
echo "<br><br>Секунҷаи ABC:";
echo "<br>HA = ".$HA;
echo "<br>HB = ".$HB;
echo "<br>HC = ".$HC;
Heights($xA, $yA, $xB, $yB, $xD, $yD, $HA, $HB, $HC);
echo "<br><br>Секунҷаи ABD:";
echo "<br>HA = ".$HA;
echo "<br>HB = ".$HB;
echo "<br>HD = ".$HC;
Heights($xB, $yB, $xC, $yC, $xD, $yD, $HA, $HB, $HC);
echo "<br><br>Секунҷаи BCD:";
echo "<br>HB = ".$HA;
echo "<br>HC = ".$HB;
echo "<br>HD = ".$HC;
?>