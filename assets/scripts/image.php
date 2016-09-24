<?php
error_reporting(E_ALL);
ini_set('display_errors', 1);
ini_set("memory_limit","256M");

$name = 'images/template.jpg';
$d = parseText();
$im = writeToImage($name, $d);
imagejpeg($im, $d['tweet_id']);

/*** spit the image out the other end ***/
function writeToImage($imagefile, $d){
    /*** make sure the file exists ***/
    if(file_exists($imagefile)){
        /*** create image ***/
        $im = @imagecreatefromjpeg($imagefile);
        /*** create the text color ***/
        $font = 'fonts/source-sans-pro.bold.ttf';
        $font2 = 'fonts/source-sans-pro.regular.ttf';
        $white = imagecolorallocate($im, 255, 255, 255);
        $grey = imagecolorallocate($im, 190, 190, 190);
        // 		Draw


        $dateTime = strtotime($d['returns']);
        $text_start_x = 70;
        $text_start_y = 275;
        $y_diff = 50;
        $x_diff = 0;
        imagettftextSp($im, 75, 0,  $text_start_x, $text_start_y, $white, $font, $d['name'], -0.00000001);
        imagettftextSp($im, 20, 0, $text_start_x + $x_diff, $text_start_y + $y_diff, $grey, $font, "RETURNING", 0);
        $num = numDays(date('Y-m-d', $dateTime));

        if($num > 0){
            $dayType = " days)";
            if($num == 1){
                $dayType = " day)";
            }
            imagettftextSp($im, 23, 0, $text_start_x+ $x_diff +150, $text_start_y + $y_diff, $white, $font2, date('D jS M', $dateTime). " (".$num.$dayType, -0.00000001);
        }
        else{
            imagettftextSp($im, 23, 0, $text_start_x+ $x_diff +150, $text_start_y + $y_diff, $white, $font2, date('D jS M', $dateTime), -0.00000001);
        }
        imagettftextSp($im, 20, 0, $text_start_x + $x_diff +500, $text_start_y + $y_diff, $grey, $font, "INJURY", 0);
        imagettftextSp($im, 23, 0, $text_start_x + $x_diff +600, $text_start_y + $y_diff, $white, $font2, $d['injury'], -0.00000001);
        if(array_key_exists('image',$d)){

            $loc  = "images/players/".str_replace(" ", "-",$d['name'].".jpg");

            /* 			Add wiki image
            $temp = save_image($d['image'], $loc);
            chmod($loc,0755);
            // 			Resize it
            $size = 140;
            $crop = crop($loc, $loc, $size);
            $player = @imagecreatefromjpeg($loc);
            list($width, $height) = getimagesize($loc);
            imagecopyresampled($im, $player, 1024 - $size - 35, 30, 0, 0, $width, $height, $width, $height);
            /*			Add circle crop
            $circle_x = 250;
            $circle_y = 175;
            $circle_loc = "images/circle.png";
            $circle = @imagecreatefrompng($circle_loc);
            list($width, $height) = getimagesize($circle_loc);
            imagecopyresampled($im, $circle, 1024 - $circle_x, 0, 0, 0, $circle_x, $circle_y, $circle_x, $circle_y); */
            //			Add logo
            $logo = "images/logos/".$d['team'].".png";

            $logo_size = 100;
            if(file_exists($logo)){
                $logo_img = @imagecreatefrompng($logo);
                list($width, $height) = getimagesize($logo);
                imagecopyresampled($im, $logo_img, 895, 25, 0, 0, $logo_size, $logo_size, $logo_size, $logo_size);
            }

        }
    }

    return $im;
}
function save_image($img,$fullpath){
    $fp = fopen ($fullpath, 'w+');
    // 	open file handle
    $ch = curl_init($img);
    curl_setopt($ch, CURLOPT_FILE, $fp);
    curl_setopt($ch, CURLOPT_SSL_VERIFYHOST, false);
    curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, false);
    curl_setopt($ch, CURLOPT_FOLLOWLOCATION, 1);
    curl_setopt($ch, CURLOPT_TIMEOUT, 1000);
    curl_setopt($ch, CURLOPT_USERAGENT, 'Mozilla/5.0');
    curl_exec($ch);
    curl_close($ch);
    fclose($fp);
}
function parseText(){
    $ret = array();
    $params = explode(",",$_GET['data']);
    $ret['name'] = urldecode($params[0]);
    $ret['injury'] = urldecode($params[1]);
    $ret['team'] = urldecode($params[2]);

    $ret['perc'] = urldecode($params[3]);
    $ret['returns'] = urldecode($params[4]);
    $ret['tweet_id']  = "output/".$params[5].".jpg";

    // 	Wiki
    $url ="https://en.wikipedia.org/w/api.php?action=query&prop=extracts|pageimages&format=json&exlimit=10&exintro=&exsentences=1&explaintext=&piprop=thumbnail&pithumbsize=400&pilimit=10&redirects=&titles=";
    $arrContextOptions=array(
    "ssl"=>array(
    "verify_peer"=>false,
    "verify_peer_name"=>false,
    ),
    );
    $full = $url.rawurlencode($ret['name']);
    //		echo "Opening ".$full;
    $json = json_decode(file_get_contents($full,false, stream_context_create($arrContextOptions)),true);
    $img = "";
    if(array_key_exists("query", $json) && array_key_exists("pages", $json['query'])){
        $sub = $json['query']['pages'];
        $keys = array_keys($sub);
        $cur = $sub[$keys[0]];
        if(array_key_exists("thumbnail", $cur) && array_key_exists("source", $cur['thumbnail'])){
            $ret['image'] = $cur['thumbnail']['source'];
        }
    }
    return $ret;
}
function crop($imgSrc, $loc, $size){
    //g	etting the image dimensions
    list($width, $height) = getimagesize($imgSrc);
    //s	aving the image into memory (for manipulation with GD Library)
    $myImage = imagecreatefromjpeg($imgSrc);
    // 	calculating the part of the image to use for thumbnail
    if ($width > $height) {
        $y = 10;
        $x = ($width - $height) / 2;
        $smallestSide = $height;
    }
    else {
        $x = -10;
        $y = 0;//($height - $width) / 2;
        $smallestSide = $width;
    }
    // 	copying the part into thumbnail
    $thumbSize = $size;
    $thumb = imagecreatetruecolor($thumbSize, $thumbSize);
    imagecopyresampled($thumb, $myImage, 0, 0, $x, $y, $thumbSize, $thumbSize, $smallestSide, $smallestSide);
    //f	inal output
    return imagejpeg($thumb, $loc);
}
function imagettftextSp($image, $size, $angle, $x, $y, $color, $font, $text, $spacing = 0)
{
    if ($spacing == 0) {
        imagettftext($image, $size, $angle, $x, $y, $color, $font, $text);
    }
    else {
        $temp_x = $x;
        $temp_y = $y;
        //t		o avoid special char problems
        $char_array = preg_split('//u',$text, -1, PREG_SPLIT_NO_EMPTY);
        foreach($char_array as $char) {
            imagettftext($image, $size, $angle, $temp_x, $temp_y, $color, $font, $char);
            $bbox = imagettfbbox($size, 0, $font, $char);
            $temp_x += cos(deg2rad($angle)) * ($spacing + ($bbox[2] - $bbox[0]));
            $temp_y -= sin(deg2rad($angle)) * ($spacing + ($bbox[2] - $bbox[0]));
        }
    }
}
function numDays($d1){
    $now = time();
    // 	or your date as well
    $your_date = strtotime($d1);
    //e	cho $now." ".$your_date;
    $datediff =  $your_date - $now;
    return floor($datediff / (60 * 60 * 24));
}
?>