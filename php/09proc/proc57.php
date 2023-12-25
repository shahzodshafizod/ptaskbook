<?php
function Leng($xA, $yA, $xB, $yB){
        return number_format(sqrt(pow(xA - xB, 2) + pow(yA - yB, 2)), 2);
}
function Perim($xA, $yA, $xB, $yB, $xC, $yC){
        $a = Leng($xA, $yA, $xB, $yB);
        $b = Leng($xB, $yB, $xC, $yC);
        $c = Leng($xA, $yA, $xC, $yC);
        $P = $a + $b + $c;
        return $P;
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
$ABC = Perim($xA, $yA, $xB, $yB, $xC, $yC);
$ABD = Perim($xA, $yA, $xB, $yB, $xD, $yD);
$ACD = Perim($xA, $yA, $xC, $yC, $xD, $yD);
echo "<br>P_ABC = ".$ABC;
echo "<br>P_ABD = ".$ABD;
echo "<br>P_ACD = ".$ACD
?>