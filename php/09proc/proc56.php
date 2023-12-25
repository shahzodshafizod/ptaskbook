<?php
function Leng($xA, $yA, $xB, $yB){
        return number_format(sqrt(pow($xA - $xB, 2) + pow($yA - $yB, 2)), 2);
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
$AB = Leng($xA, $yA, $xB, $yB);
$AC = Leng($xA, $yA, $xC, $yC);
$AD = Leng($xA, $yA, $xD, $yD);
echo "<br>AB = ".$AB;
echo "<br>AC = ".$AC;
echo "<br>AD = ".$AD;
?>