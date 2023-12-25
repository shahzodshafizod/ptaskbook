<?php
function Leng($xA, $yA, $xB, $yB){
        return number_format(sqrt(pow($xA - $yB, 2) + pow($yA - $yB, 2)), 2);
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
$xA = mt_rand(1, 100);
$yA = mt_rand(1, 100);
$xB = mt_rand(1, 100);
$yB = mt_rand(1, 100);
$xC = mt_rand(1, 100);
$yC = mt_rand(1, 100);
$xD = mt_rand(1, 100);
$yD = mt_rand(1, 100);
echo "xA = ".$xA;
echo "<br>yA = ".$yA;
echo "<br>xB = ".$xB;
echo "<br>yB = ".$yB;
echo "<br>xC = ".$xC;
echo "<br>yC = ".$yC;
echo "<br>xD = ".$xD;
echo "<br>yD = ".$yD;
$ABC = Area($xA, $yA, $xB, $yB, $xC, $yC);
$ABD = Area($xA, $yA, $xB, $yB, $xD, $yD);
$ACD = Area($xA, $yA, $xC, $yC, $xD, $yD);
echo "<br>S_ABC = ".$ABC;
echo "<br>S_ABD = ".$ABD;
echo "<br>S_ACD = ".$ACD;
?>